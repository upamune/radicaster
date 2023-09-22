package http

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/upamune/podcast-server/podcast"
)

func NewHTTPHandler(
	podcaster *podcast.Podcaster,
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
		c.Blob(http.StatusOK, "application/xml", []byte(podcaster.GetFeed()))
		return nil
	})
	e.GET("/sync", func(c echo.Context) error {
		podcaster.Sync()
		c.String(http.StatusOK, "")
		return nil
	})

	e.GET("/rss.xml", func(c echo.Context) error {
		c.Blob(http.StatusOK, "application/xml", []byte(podcaster.GetFeed()))
		return nil
	})

	e.Static("/static", targetDir)

	return e, nil
}
