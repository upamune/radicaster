package http

import (
	"bytes"
	"crypto/subtle"
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/upamune/radicaster/config"
	"github.com/upamune/radicaster/podcast"
	"github.com/upamune/radicaster/record"
)

//go:embed views
var views embed.FS

func NewHTTPHandler(
	logger zerolog.Logger,
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
		return c.Blob(http.StatusOK, "application/xml", []byte(podcaster.GetFeed()))
	})
	e.GET("/sync", func(c echo.Context) error {
		if err := podcaster.Sync(); err != nil {
			return c.String(http.StatusInternalServerError, "")
		}
		return c.String(http.StatusOK, "")
	})

	e.GET("/rss.xml", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "application/xml", []byte(podcaster.GetFeed()))
	})

	t := template.Must(template.New("config.html.tmpl").
		Funcs(template.FuncMap{
			"inc": func(i int) int {
				return i + 1
			},
		}).
		ParseFS(views, "views/config.html.tmpl"))
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

		var buf bytes.Buffer
		if err := t.Execute(
			&buf,
			map[string]interface{}{"Programs": config.Programs},
		); err != nil {
			return c.String(http.StatusInternalServerError, "")
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
