package main

import (
	"context"
	"flag"
	"fmt"
	httpstd "net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/upamune/podcast-server/http"
	"github.com/upamune/podcast-server/podcast"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	baseURL := flag.String("baseurl", "", "base URL of server")
	targetDir := flag.String("targetdir", "", "audio target directory")
	basicAuth := flag.String("basicauth", "", "basic auth for HTTP server")
	flag.Parse()

	if baseURL == nil || *baseURL == "" {
		fmt.Fprintf(os.Stderr, "-baseurl is required")
		return 1
	}

	if targetDir == nil || *targetDir == "" {
		fmt.Fprintf(os.Stderr, "-targetdir is required")
		return 1
	}

	logger := zerolog.New(os.Stderr)

	now := time.Now()
	podcaster := podcast.NewPodcaster(
		logger,
		*baseURL,
		*targetDir,
		"Radiko",
		*baseURL,
		"Radiko",
		&now,
	)

	if err := podcaster.Sync(); err != nil {
		logger.Error().Err(err).Msg("failed to initial sync")
		return 1
	}
	handler, err := http.NewHTTPHandler(podcaster, *targetDir, *basicAuth)
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

	if err := server.ListenAndServe(); err != nil {
		logger.Error().Err(err).Msg("failed to listen and serve")
		return 1
	}

	return 0
}
