package routes

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"goblin/config"
	mw "goblin/pkg/middleware"
	"goblin/pkg/services"
)

func BuildRouter(c *services.Container) {
	c.Web.HideBanner = true

	// Enable cache control for static files.
	// NOTE: We need to use funcmap.File() to append a cache key to the URL in
	// order to break the cache after each server restart.
	c.Web.Group("", mw.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Non-static routes
	g := c.Web.Group("")

	// Force HTTPS if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.RecoverWithConfig(echomw.RecoverConfig{
			StackSize: 1 << 10, // 1 KB
			LogLevel:  log.ERROR,
		}),
		echomw.Secure(),
		echomw.RequestID(),
		// RquestID middleware must be called before LogRequestID.
		mw.LogRequestID(),
		// TODO: See https://github.com/labstack/echo/issues/1223
		mw.Brotli(),
		// TODO: Replace with zap logger
		// https://betterstack.com/community/guides/logging/go/zap/
		echomw.Logger(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup:    "form:csrf",
			CookieSameSite: http.SameSiteDefaultMode,
		}),
	)

	g.GET("/", func(c echo.Context) error {
		return c.String(200, "{ok}")
	})
}
