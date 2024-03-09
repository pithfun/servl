package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/tiny-blob/tinyblob/pkg/htmx"
	"github.com/tiny-blob/tinyblob/templates"
)

// Page consists of all data that will be used to render a page response for a
// given controller.
// While it’s not required for a controller to render a Page on a route, this is
// the common data object that will be passed to the templates, making it easy
// for all controllers to share functionality both on the back and frontend. The
// page can be expanded to include anything else your app wants to support.
// Methods on this page also then become available in the templates, which can
// be more useful than the funcmap if your methods require data store in the
// page, such as the context.
type Page struct {
	// AppName stores the name of the application. If omitted, the configuration
	// value will be used.
	AppName string
	// Context stores the request context.
	Context echo.Context
	// CSRF stores the CSRF token for the given request.
	// This will only be populated if the CSRF middleware is in effect for the
	// given request. If this is populated, all forms must include this value
	// otherwise the requests will be rejected.
	CSRF string
	// Headers stores a list of HTTP headers to be set on the response.
	Headers map[string]string
	// IsAuth stores whether or not the user is authenticated.
	IsAuth bool
	// IsHome stores whether or not the requested page is the home page or not.
	IsHome bool
	// Layout stores the name of the layout base template file which will be used
	// when the page is rendered.
	// This should match a template file located within the layouts directory
	// inside the templates directory. The template extension should not be
	// included in this value.
	Layout templates.Layout
	// Name stores the name of the page as well as the name of the template file
	// which will be used to render the content portion of the layout template.
	// This should match a template file located within the pages directory inside
	// the templates directory. The template extension should not be included in
	// this value.
	Name templates.Page
	// RequestID stores the ID of the given request.
	// This will only be populated if the request ID middleware is in effect for the given request.
	RequestID string
	// StatusCode stores the HTTP status code to be returned.
	StatusCode int
	// ToURL is a function to convert a route name and optional route parameters to a URL
	ToURL func(name string, params ...any) string
	Path  string
	// URL stores the URL of the current request
	URL string
	// Title stores the title of the page.
	Title string
	// Cache stores values for caching the response of this page.
	Cache struct {
		// Enabled dictates if the response of this page should be cached.
		Enabled bool
		// Expiration stores the amount of time that the cache entry should live for before expiring.
		// If omitted, the configuration value will be used.
		Expiration time.Duration
		// Tags stores a list of tags to apply to the cache entry.
		// These are useful when invalidating cache for dynamic events such as entity operations.
		Tags []string
	}
	// HTMX stores the Request and Response values from htmx.
	HTMX struct {
		Request  htmx.Request
		Response *htmx.Response
	}
	// Metatags stores metatag values
	Metatags struct {
		// Description stores the description metatag value.
		Description string
		// Keywords stores the keywords metatag values.
		Keywords []string
	}
}

// NewPage creates and initiatizes a new Page for a given request context
func NewPage(ctx echo.Context) Page {
	p := Page{
		Context:    ctx,
		Path:       ctx.Request().URL.Path,
		Headers:    make(map[string]string),
		RequestID:  ctx.Response().Header().Get(echo.HeaderXRequestID),
		StatusCode: http.StatusOK,
		ToURL:      ctx.Echo().Reverse,
		URL:        ctx.Request().URL.String(),
	}

	p.IsHome = p.Path == "/"

	if csrf := ctx.Get(echomw.DefaultCSRFConfig.ContextKey); csrf != nil {
		p.CSRF = csrf.(string)
	}

	// TODO: Authentication

	p.HTMX.Request = htmx.GetRequest(ctx)

	return p
}
