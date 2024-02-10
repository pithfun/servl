package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/makomarket/mako/pkg/context"
	"github.com/makomarket/mako/pkg/middleware"
	"github.com/makomarket/mako/pkg/services"
)

// Controller provides base functionality and dependencies to routes.
// A controller is embedded in each individual route struct and is used by the
// router to inject the container so that each route has access to the services
// within the container.
type Controller struct {
	// Container stores a services container which contains contains dependencies.
	Container *services.Container
}

// NewController creates a new controller.
func NewController(c *services.Container) Controller {
	return Controller{
		Container: c,
	}
}

// RenderPage renders a page as a HTTP response.
func (c *Controller) RenderPage(ctx echo.Context, page Page) error {
	var buf *bytes.Buffer
	var err error

	// Page name is required
	if page.Name == "" {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"page name is required to render template",
		)
	}

	// Use the app name in configuration if a value was not set
	if page.AppName == "" {
		page.AppName = c.Container.Config.App.Name
	}

	// Check if this is an HTMX non-boosted request which indicates that only
	// partial content should be rendered.
	if page.HTMX.Request.Enabled && !page.HTMX.Request.Boosted {
		// Parse and execute the templates only for the content portion of the page.
		// The templates used for this partial request will be:
		// 1. The base HTMX template which omits the layout and only includes the content template.
		// 2. The content template specified in Page.Name.
		// 3. All templates within the components directory.
		// Also included is the function map provided by the funcmap package.
		buf, err = c.Container.TemplateRenderer.
			Parse().
			Group("page:htmx").
			Key(page.Name).
			Base("htmx").
			Files(
				"htmx",
				fmt.Sprintf("pages/%s", page.Name),
			).
			Directories("components").
			Execute(page)
	} else {
		// Parse and execute the templates for the page.
		// The templates used for the page will be:
		// 1. The layout/base template specified in Page.Layout
		// 2. The content template specified in Page.Name
		// 3. All templates within the components directory
		// Also included is the function map provided by the funcmap package
		buf, err = c.Container.TemplateRenderer.
			Parse().
			Group("page").
			Key(page.Name).
			Base(page.Layout).
			Files(
				fmt.Sprintf("layouts/%s", page.Layout),
				fmt.Sprintf("pages/%s", page.Name),
			).
			Directories("components").
			Execute(page)
	}

	if err != nil {
		return c.Fail(err, "failed to parse and execute template")
	}

	// Set the status code.
	ctx.Response().Status = page.StatusCode

	// Set response headers.
	for k, v := range page.Headers {
		ctx.Response().Header().Set(k, v)
	}

	// Apply the HTMX response, if one
	if page.HTMX.Response != nil {
		page.HTMX.Response.Apply(ctx)
	}

	// If caching is enabled, cache the page.
	c.cachePage(ctx, page, buf)

	return ctx.HTMLBlob(ctx.Response().Status, buf.Bytes())
}

// cachePage caches the HTML for a given Page if the Page has cache enabled.
func (c *Controller) cachePage(ctx echo.Context, page Page, html *bytes.Buffer) {
	if !page.Cache.Enabled || page.IsAuth {
		return
	}

	// Check that this is a valid buffer
	if html == nil || html.Len() == 0 {
		return
	}

	// If we did not specify an expiration, default to the configuration expiration.
	if page.Cache.Expiration == 0 {
		page.Cache.Expiration = c.Container.Config.Cache.Expiration.Page
	}

	// Extract the headers
	headers := make(map[string]string)
	for k, v := range ctx.Response().Header() {
		headers[k] = v[0]
	}

	// The request URL is used as the cache key so the middleware can serve the
	// cached page on matching requests.
	key := ctx.Request().URL.String()
	cp := middleware.CachedPage{
		Headers:    headers,
		HTML:       html.Bytes(),
		StatusCode: ctx.Response().Status,
		URL:        key,
	}

	err := c.Container.Cache.
		Set().
		Group(middleware.CachedPageGroup).
		Key(key).
		Tags(page.Cache.Tags...).
		Expiration(page.Cache.Expiration).
		Data(cp).
		Save(ctx.Request().Context())

	switch {
	case err == nil:
		ctx.Logger().Infof("cached page with key: %v", key)
	case !context.IsCanceledError(err):
		ctx.Logger().Errorf("failed to cache page: %v", err)
	}
}

// Fail is a helper to fail a request by returning a 500 error and logging the error
func (c *Controller) Fail(err error, log string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%s: %v", log, err))
}
