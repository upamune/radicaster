// Copy from https://github.com/yyoshiki41/radigo/blob/7eb0cce/ffmpeg.go
package ffmpeg

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
)

type ffmpeg struct {
	*exec.Cmd
}

func newFfmpeg(ctx context.Context) (*ffmpeg, error) {
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	return &ffmpeg{exec.CommandContext(
		ctx,
		cmdPath,
	)}, nil
}

func (f *ffmpeg) setDir(dir string) {
	f.Dir = dir
}

func (f *ffmpeg) setArgs(args ...string) {
	f.Args = append(f.Args, args...)
}

func (f *ffmpeg) setInput(input string) {
	f.setArgs("-i", input)
}

func (f *ffmpeg) run(output string) error {
	f.setArgs(output)
	return f.Run()
}

func (f *ffmpeg) start(output string) error {
	f.setArgs(output)
	return f.Start()
}

func (f *ffmpeg) wait() error {
	return f.Wait()
}

func (f *ffmpeg) stdinPipe() (io.WriteCloser, error) {
	return f.StdinPipe()
}

func (f *ffmpeg) stderrPipe() (io.ReadCloser, error) {
	return f.StderrPipe()
}

// ConcatAACFilesFromList concatenates files from the list of resources.
func ConcatAACFilesFromList(ctx context.Context, logger zerolog.Logger, resourcesDir string) (string, error) {
	files, err := os.ReadDir(resourcesDir)
	if err != nil {
		return "", err
	}

	allFilePaths := []string{}
	for _, f := range files {
		p := filepath.Join(resourcesDir, f.Name())
		allFilePaths = append(allFilePaths, p)
	}
	concatedFile := filepath.Join(resourcesDir, "concated.aac")
	if err := ConcatAACFilesAll(ctx, logger, allFilePaths, resourcesDir, concatedFile); err != nil {
		return "", err
	}

	return concatedFile, nil
}

// ConcatAACFilesAll concatenate files of the same type.
func ConcatAACFilesAll(ctx context.Context, logger zerolog.Logger, files []string, resourcesDir string, output string) error {
	// input is a path to a file which lists all the aac files.
	// it may include a lot of aac file and exceed max number of file descriptor.
	oneConcatNum := 100
	if len(files) > oneConcatNum {
		reducedFiles := files[:oneConcatNum]
		restFiles := files[oneConcatNum:]
		// reducedFiles -> reducedFiles[0]
		tmpOutputFile, err := os.CreateTemp(resourcesDir, "tmp-concatenated-*.aac")
		if err != nil {
			return errors.Wrap(err, "failed to create a temporary file")
		}
		defer os.Remove(tmpOutputFile.Name())

		if err := ConcatAACFiles(ctx, logger, reducedFiles, resourcesDir, tmpOutputFile.Name()); err != nil {
			return errors.Wrap(err, "failed to concatenate aac files")
		}
		if err := ConcatAACFilesAll(ctx, logger, append([]string{tmpOutputFile.Name()}, restFiles...), resourcesDir, output); err != nil {
			return errors.Wrap(err, "failed to concatenate aac files")
		}

		return nil
	} else {
		return ConcatAACFiles(ctx, logger, files, resourcesDir, output)
	}
}

func ConcatAACFiles(ctx context.Context, logger zerolog.Logger, input []string, resourcesDir string, output string) error {
	listFile, err := os.CreateTemp(resourcesDir, "aac_resources")
	if err != nil {
		return err
	}
	defer os.Remove(listFile.Name())

	for _, f := range input {
		p := fmt.Sprintf("file '%s'\n", f)
		if _, err := listFile.WriteString(p); err != nil {
			return errors.Wrap(err, "failed to write aac file path to the list file")
		}
	}

	f, err := newFfmpeg(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to create ffmpeg command")
	}

	f.setArgs(
		"-nostdin",
		"-f", "concat",
		"-safe", "0",
		"-y",
	)
	f.setInput(listFile.Name())
	f.setArgs("-c", "copy")

	logger.Debug().
		Str("label", "ffmpeg").
		Str("command", f.String()).
		Msg("concatenate aac files by ffmpeg")

	defer func() {
		for _, f := range input {
			os.Remove(f)
		}
	}()

	stdErrPipe, err := f.stderrPipe()
	if err != nil {
		return errors.Wrap(err, "failed to get stderr pipe")
	}

	f.setArgs(output)
	if err := f.Start(); err != nil {
		return errors.Wrap(err, "failed to start ffmpeg")
	}

	b, err := io.ReadAll(stdErrPipe)
	if err != nil {
		return errors.Wrap(err, "failed to read stderr")
	}

	logger.Debug().
		Str("label", "ffmpeg").
		Str("stderr", string(b)).
		Msg("concatenate aac files by ffmpeg")

	if err := f.Wait(); err != nil {
		return errors.Wrap(err, "failed to wait ffmpeg")
	}

	return nil
}
