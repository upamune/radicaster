package main

import (
	"flag"
	"fmt"
	httpstd "net/http"
	"os"
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

	httpstd.ListenAndServe(":3333", handler)

	return 0
}
