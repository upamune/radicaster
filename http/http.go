package http

import (
	"bytes"
	"context"
	"crypto/subtle"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/goccy/go-yaml"
	"github.com/jellydator/ttlcache/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"github.com/upamune/radicaster/config"
	"github.com/upamune/radicaster/podcast"
	"github.com/upamune/radicaster/radikoutil"
	"github.com/upamune/radicaster/record"
	"github.com/yyoshiki41/go-radiko"
)

//go:embed views
var views embed.FS

func NewHTTPHandler(
	logger zerolog.Logger,
	version, revision string,
	podcaster *podcast.Podcaster,
	recorder *record.Recorder,
	targetDir string,
	basicAuth string,
) (http.Handler, error) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(noCacheMiddleware)
	if ss := strings.Split(basicAuth, ":"); len(ss) == 2 {
		e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
			Skipper: func(c echo.Context) bool {
				path := c.Request().URL.Path
				// NOTE: 静的ファイルの配信はベージック認証をスキップする
				return strings.HasPrefix(path, "/static")
			},
			Validator: func(username, password string, c echo.Context) (bool, error) {
				if subtle.ConstantTimeCompare([]byte(username), []byte(ss[0])) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(ss[1])) == 1 {
					return true, nil
				}
				return false, nil
			},
			Realm: "Restricted",
		}))
	}

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/config")
	})

	e.GET("/sync", func(c echo.Context) error {
		if err := podcaster.Sync(); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}
		return c.String(http.StatusOK, "")
	})

	e.GET("/rss.xml", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "application/xml", []byte(podcaster.GetDefaultFeed()))
	})

	e.GET("/zenroku/:program_path/rss.xml", func(c echo.Context) error {
		programPath := c.Param("program_path")
		p := path.Join("zenroku", programPath)
		feed, ok := podcaster.GetFeed(p)
		if !ok {
			return c.String(http.StatusNotFound, "")
		}
		return c.Blob(
			http.StatusOK,
			"application/xml",
			[]byte(feed),
		)
	})

	e.GET("/:program_path/rss.xml", func(c echo.Context) error {
		programPath := c.Param("program_path")
		feed, ok := podcaster.GetFeed(programPath)
		if !ok {
			return c.String(http.StatusNotFound, "")
		}
		return c.Blob(
			http.StatusOK,
			"application/xml",
			[]byte(feed),
		)
	})

	t, err := template.New("config.html.tmpl").
		Funcs(template.FuncMap{
			"inc": func(i int) int {
				return i + 1
			},
			"sliceToStr": func(s []string) string {
				return strings.Join(s, ", ")
			},
			"isEnabled": func(id string, enableStationIDs []string) bool {
				id = strings.ToLower(id)
				for _, enableStationID := range enableStationIDs {
					if id == strings.ToLower(enableStationID) {
						return true
					}
				}
				return false
			},
		}).
		ParseFS(views, "views/config.html.tmpl")
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse template")
	}

	radikoCache := ttlcache.New[string, radiko.Stations](
		ttlcache.WithTTL[string, radiko.Stations](24 * time.Hour),
	)
	e.GET("/config", func(c echo.Context) error {
		config := recorder.Config()

		acceptHeader := c.Request().Header.Get("Accept")
		if acceptHeader == "application/json" || acceptHeader == "json" {
			return c.JSON(http.StatusOK, config)
		}

		if acceptHeader == "application/yaml" || acceptHeader == "yaml" {
			var buf bytes.Buffer
			if err := yaml.
				NewEncoder(&buf, yaml.Indent(2)).
				Encode(config); err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.Blob(http.StatusOK, "application/yaml", buf.Bytes())
		}

		type zenrokuStation struct {
			ImageURL string
			ID       string
			Name     string
			Path     string
			Enabled  bool
		}

		var zenrokuStations []zenrokuStation

		if config.Zenroku.Enable {
			ctx := context.Background()

			const radikoCacheKey = "stations"
			var stations radiko.Stations
			if radikoCache.Has(radikoCacheKey) {
				item := radikoCache.Get(radikoCacheKey)
				logger.Debug().
					Str("cache_key", radikoCacheKey).
					Bool("is_expired", item.IsExpired()).
					Time("expires_at", item.ExpiresAt()).
					Msg("radiko cache hit")
				stations = item.Value()
			} else {
				logger.Debug().
					Str("cache_key", radikoCacheKey).
					Msg("radiko cache no-hit")
				client, err := radikoutil.NewClient(ctx)
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
				stations, err = client.GetStations(ctx, time.Now())
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
				radikoCache.Set(radikoCacheKey, stations, 24*time.Hour)
			}

			enableStationIDMap := lo.Associate(config.Zenroku.EnableStationIDs, func(stationID string) (string, struct{}) {
				return strings.ToLower(stationID), struct{}{}
			})

			for _, station := range stations {
				station := station
				stationID := strings.ToLower(station.ID)
				_, enabled := enableStationIDMap[stationID]
				zenrokuStations = append(zenrokuStations, zenrokuStation{
					ImageURL: config.Zenroku.Stations[stationID].ImageURL,
					ID:       station.ID,
					Name:     station.Name,
					Path:     fmt.Sprintf("/zenroku/%s", stationID),
					Enabled:  enabled,
				})
			}
		}

		var buf bytes.Buffer
		if err := t.Execute(
			&buf,
			map[string]interface{}{
				"Programs":        config.Programs,
				"Zenroku":         config.Zenroku,
				"ZenrokuStations": zenrokuStations,
				"Version":         version,
				"Revision":        revision,
			},
		); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.HTML(http.StatusOK, buf.String())
	})

	e.PUT("/config", func(ctx echo.Context) error {
		var c config.Config

		if ctx.Request().Header.Get("Content-Type") == "application/yaml" {
			body := ctx.Request().Body
			defer body.Close()
			var err error
			c, err = config.Parse(body)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
		} else {
			if err := ctx.Bind(&c); err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
		}

		if err := c.Validate(); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		updatedConfig, err := recorder.RefreshConfig(c)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		acceptHeader := ctx.Request().Header.Get("Accept")
		if acceptHeader == "application/yaml" || acceptHeader == "yaml" {
			var buf bytes.Buffer
			if err := yaml.
				NewEncoder(&buf, yaml.Indent(2)).
				Encode(updatedConfig); err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}
			return ctx.Blob(http.StatusOK, "application/yaml", buf.Bytes())
		}

		return ctx.JSON(http.StatusOK, updatedConfig)
	})

	e.Static("/static", targetDir)

	return e, nil
}
