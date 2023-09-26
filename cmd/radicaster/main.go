package main

import (
	"context"
	"flag"
	"fmt"
	httpstd "net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/upamune/radicaster/config"
	"github.com/upamune/radicaster/http"
	"github.com/upamune/radicaster/podcast"
	"github.com/upamune/radicaster/radikoutil"
	"github.com/upamune/radicaster/record"
)

var (
	Version  string
	Revision string
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	baseURL := flag.String("baseurl", "http://localhost:3333", "base URL of server")
	targetDir := flag.String("targetdir", "./output", "audio target directory")
	basicAuth := flag.String("basicauth", "", "Basic認証のための ':' で区切られたユーザー名とパスワード")
	programConfig := flag.String("config", "./radicast.yaml", "path for config")
	programConfigURL := flag.String("configurl", "", "url for config")
	podcastImageURL := flag.String("podcastimageurl", "", "url for podcast image")
	flag.Parse()

	if baseURL == nil || *baseURL == "" {
		fmt.Fprintf(os.Stderr, "-baseurl is required")
		return 1
	}

	if targetDir == nil || *targetDir == "" {
		fmt.Fprintf(os.Stderr, "-targetdir is required")
		return 1
	}

	logger := zerolog.New(os.Stderr).
		With().
		Str("version", Version).
		Str("revision", Revision).
		Logger()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error().Err(err).Msg("failed to create file watcher")
		return 1
	}
	defer watcher.Close()

	now := time.Now()
	podcaster := podcast.NewPodcaster(
		logger,
		*baseURL,
		*targetDir,
		"Radicaster",
		*baseURL,
		"Radicaster",
		&now,
		*podcastImageURL,
	)

	if err := podcaster.Sync(); err != nil {
		logger.Error().Err(err).Msg("failed to initial sync")
		return 1
	}

	go func() {
		logger := logger.With().Str("component", "file_watcher").Logger()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logger.Info().
					Str("event_name", event.Name).
					Str("event_operator", event.Op.String()).
					Msg("got a file event")

				if event.Has(fsnotify.Chmod) {
					continue
				}
				if err := podcaster.Sync(); err != nil {
					logger.Error().Err(err).Msg("failed to sync")
					continue
				}
				logger.Info().Msg("synced")
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error().Err(err).Msg("got a file watcher error")
			}
		}
	}()
	if err := watcher.Add(*targetDir); err != nil {
		logger.Error().Err(err).Msg("failed to add file watcher")
		return 1
	}
	logger.Info().Str("target_dir", *targetDir).Msg("added file watcher for sync")

	initConfig, err := config.Init(programConfig, programConfigURL)
	if err != nil {
		logger.Error().Err(err).Msg("failed to init config")
		return 1
	}

	ctx := context.Background()
	radikoClient, err := radikoutil.NewClient(ctx)
	recorder, err := record.NewRecorder(logger, radikoClient, *targetDir, initConfig)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create recorder")
		return 1
	}

	handler, err := http.NewHTTPHandler(logger, podcaster, recorder, *targetDir, *basicAuth)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create HTTP handler")
		return 1
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := &httpstd.Server{Addr: ":3333", Handler: handler}

	go func() {
		<-ctx.Done()
		logger.Info().Msg("shutting down server in 60 seconds")
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Error().Err(err).Msg("failed to shutdown server")
			return
		}
	}()

	logger.Info().Str("base_url", *baseURL).Msg("http server is starting...")
	if err := server.ListenAndServe(); err != nil {
		logger.Error().Err(err).Msg("failed to listen and serve")
		return 1
	}

	return 0
}
