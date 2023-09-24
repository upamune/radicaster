package http

import (
	"bytes"
	"crypto/subtle"
	"html/template"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
	"github.com/upamune/podcast-server/config"
	"github.com/upamune/podcast-server/podcast"
	"github.com/upamune/podcast-server/record"
)

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
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(ss[0])) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(ss[1])) == 1 {
				return true, nil
			}
			return false, nil
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
		ParseFiles("http/views/config.html.tmpl"))
	e.GET("/config", func(c echo.Context) error {
		config := recorder.Config()

		if c.QueryParams().Get("format") == "json" {
			return c.JSON(http.StatusOK, config)
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

		return ctx.JSON(http.StatusOK, updatedConfig)
	})

	e.Static("/static", targetDir)

	return e, nil
}
