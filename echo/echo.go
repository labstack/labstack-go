package labstack

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/labstack-go/cube"
)

func Cube(key string, options cube.Options) echo.MiddlewareFunc {
	c := cube.New(key, options)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			// Start
			req := ctx.Request()
			res := ctx.Response()
			bytesIn, err := strconv.ParseInt(req.Header.Get("Content-Length"), 10, 64)
			r := &cube.Request{
				ID:        req.Header.Get(echo.HeaderXRequestID),
				Host:      req.Host,
				Path:      req.URL.Path,
				Method:    req.Method,
				BytesIn:   bytesIn,
				RemoteIP:  ctx.RealIP(),
				ClientID:  ctx.RealIP(),
				UserAgent: req.UserAgent(),
			}
			c.Start(r)

			// Next
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}

			// Stop
			if r.ID == "" {
				r.ID = res.Header().Get(echo.HeaderXRequestID)
			}
			r.Status = res.Status
			r.BytesOut = res.Size
			c.Stop(r)

			return nil
		}
	}
}
