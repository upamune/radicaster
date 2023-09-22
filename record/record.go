package record

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/go-co-op/gocron"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sourcegraph/conc/pool"
	"github.com/upamune/podcast-server/config"
	"github.com/yyoshiki41/go-radiko"
	"github.com/yyoshiki41/radigo"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

type Recorder struct {
	httpClient *retryablehttp.Client
	client     *radiko.Client

	targetDir string

	scheduler struct {
		sync.RWMutex
		*gocron.Scheduler
	}

	config struct {
		sync.RWMutex
		config.Config
	}
}

func NewRecorder(client *radiko.Client, targetDir string, initConfig config.Config) (*Recorder, error) {
	r := &Recorder{
		client:     client,
		httpClient: retryablehttp.NewClient(),
		targetDir:  targetDir,
	}
	r.config.Config = initConfig

	if err := r.restartScheduler(); err != nil {
		return nil, errors.Wrap(err, "failed to update scheduler")
	}

	return r, nil
}

func parseTime(s string) (time.Time, error) {
	const layout = "200601021504"
	t, err := time.ParseInLocation(layout, s, jst)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to parse time")
	}
	return t, nil
}

func (r *Recorder) Record(p config.Program) error {
	ctx := context.Background()

	from, err := parseTime(p.Start)
	if err != nil {
		return errors.Wrap(err, "failed to parse from")
	}

	program, err := r.client.GetProgramByStartTime(ctx, p.StationID, from)
	if err != nil {
		return errors.Wrap(err, "failed to get program")
	}

	uri, err := r.client.TimeshiftPlaylistM3U8(ctx, p.StationID, from)
	if err != nil {
		return errors.Wrap(err, "failed to get m3u8")
	}

	chunkURLs, err := radiko.GetChunklistFromM3U8(uri)
	if err != nil {
		return errors.Wrap(err, "failed to get chunklist")
	}

	aacDir := os.TempDir()
	defer os.RemoveAll(aacDir)

	if err := r.bulkDownload(chunkURLs, aacDir); err != nil {
		return errors.Wrap(err, "failed to download aac files")
	}

	// TODO: retry
	concatedFile, err := radigo.ConcatAACFilesFromList(ctx, aacDir)
	if err != nil {
		return errors.Wrap(err, "failed to concat aac files")
	}

	fileName := fmt.Sprintf(
		"%s_%s.mp3",
		program.Title,
		from.Format("2006年01月02日"),
	)

	output := filepath.Join(r.targetDir, fileName)
	// TODO: retry
	if err := radigo.ConvertAACtoMP3(ctx, concatedFile, output); err != nil {
		return errors.Wrap(err, "failed to convert aac to mp3")
	}
	return nil
}

func (r *Recorder) bulkDownload(urls []string, output string) error {
	p := pool.New().WithErrors()

	for i, url := range urls {
		i, url := i, url
		p.Go(func() error {
			if err := r.download(url, output); err != nil {
				return errors.Wrapf(err, "failed to download %d", i)
			}
			return nil
		})
	}
	if err := p.Wait(); err != nil {
		return errors.Wrap(err, "failed to download aac files")
	}
	return nil
}

func (r *Recorder) download(link, output string) error {
	resp, err := r.httpClient.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, fileName := filepath.Split(link)
	file, err := os.Create(filepath.Join(output, fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}

func (r *Recorder) restartScheduler() error {
	s := gocron.NewScheduler(jst)

	r.config.RLock()
	defer r.config.RUnlock()
	for _, p := range r.config.Config.Programs {
		if _, err := s.Cron(p.Cron).Do(r.Record, p); err != nil {
			return errors.Wrap(err, "failed to set cron")
		}
	}

	r.scheduler.Lock()
	defer r.scheduler.Unlock()
	if s := r.scheduler.Scheduler; s != nil {
		s.Stop()
	}
	r.scheduler.Scheduler = s
	r.scheduler.Scheduler.StartAsync()
	return nil
}

func (r *Recorder) RefreshConfig(configURL string) error {
	resp, err := http.Get(configURL)
	if err != nil {
		return errors.Wrap(err, "failed to get config via URL")
	}
	defer resp.Body.Close()

	config, err := config.Parse(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to parse config")
	}

	r.config.Lock()
	r.config.Config = config
	r.config.Unlock()

	if err := r.restartScheduler(); err != nil {
		return errors.Wrap(err, "failed to update scheduler")
	}

	return nil
}
