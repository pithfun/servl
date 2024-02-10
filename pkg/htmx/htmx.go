package htmx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTMX Headers (https://htmx.org/reference/#response_headers)
const (
	// Request headers
	HeaderBoosted               = "HX-Boosted"
	HeaderCurrentURL            = "HX-Current-URL"
	HeaderHistoryRestoreRequest = "HX-History-Restore-Request"
	HeaderPrompt                = "HX-Prompt"
	HeaderRequest               = "HX-Request"
	HeaderTarget                = "HX-Target"
	HeaderTriggerName           = "HX-Trigger-Name"
	HeaderTrigger               = "HX-Trigger"
	HeaderLocation              = "HX-Location"
	HeaderPushURL               = "HX-Push-Url"
	HeaderRedirect              = "HX-Redirect"
	HeaderRefresh               = "HX-Refresh"
	HeaderTriggerAfterSwap      = "HX-Trigger-After-Swap"
	HeaderTriggerAfterSettle    = "HX-Trigger-After-Settle"
)

type (
	Request struct {
		Boosted     bool
		Enabled     bool
		Trigger     string
		TriggerName string
		Target      string
		Prompt      string
	}

	// Response contains data that the server can communicate back to HTMX
	Response struct {
		Push               string
		Location           string
		Redirect           string
		Refresh            bool
		Trigger            string
		TriggerAfterSwap   string
		TriggerAfterSettle string
		NoContent          bool
	}
)

// GetRequest extracts HTMX data from the request
func GetRequest(ctx echo.Context) Request {
	return Request{
		Enabled:     ctx.Request().Header.Get(HeaderRequest) == "true",
		Boosted:     ctx.Request().Header.Get(HeaderBoosted) == "true",
		Trigger:     ctx.Request().Header.Get(HeaderTrigger),
		TriggerName: ctx.Request().Header.Get(HeaderTriggerName),
		Target:      ctx.Request().Header.Get(HeaderTarget),
		Prompt:      ctx.Request().Header.Get(HeaderPrompt),
	}
}

// Apply applies data from a `Response` to a server response.
func (r Response) Apply(ctx echo.Context) {
	if r.Location != "" {
		ctx.Response().Header().Set(HeaderLocation, r.Location)
	}
	if r.Push != "" {
		ctx.Response().Header().Set(HeaderPushURL, r.Push)
	}
	if r.Redirect != "" {
		ctx.Response().Header().Set(HeaderRedirect, r.Redirect)
	}
	if r.Refresh {
		ctx.Response().Header().Set(HeaderRefresh, "true")
	}
	if r.Trigger != "" {
		ctx.Response().Header().Set(HeaderTrigger, r.Trigger)
	}
	if r.TriggerAfterSwap != "" {
		ctx.Response().Header().Set(HeaderTriggerAfterSwap, r.TriggerAfterSwap)
	}
	if r.TriggerAfterSettle != "" {
		ctx.Response().Header().Set(HeaderTriggerAfterSettle, r.TriggerAfterSettle)
	}
	if r.NoContent {
		ctx.Response().Status = http.StatusNoContent
	}
}
