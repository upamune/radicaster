package http

import (
	"time"

	"github.com/labstack/echo/v4"
)

var (
	// Unix epoch time
	epoch = time.Unix(0, 0).Format(time.RFC1123)

	// Taken from https://github.com/mytrile/nocache
	noCacheHeaders = map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, private, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
	}
	etagHeaders = []string{
		"ETag",
		"If-Modified-Since",
		"If-Match",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
	}
)

func noCacheMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := c.Request()
		// Delete any ETag headers that may have been set
		for _, v := range etagHeaders {
			if req.Header.Get(v) != "" {
				req.Header.Del(v)
			}
		}

		// Set our NoCache headers
		res := c.Response()
		for k, v := range noCacheHeaders {
			res.Header().Set(k, v)
		}
		return next(c)
	}
}
