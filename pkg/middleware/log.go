package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// The HTTP X-Request-ID request header is used to trace individual HTTP
// requests from the client to the server and back again. It allows the client
// and server to correlate each HTTP request.
func LogRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)
			format := `{"time":"${time_rfc3339_nano}","id":"%s","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`
			c.Logger().SetHeader(fmt.Sprintf(format, reqID))
			return next(c)
		}
	}
}
