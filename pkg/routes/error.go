package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/tiny-blob/tinyblob/pkg/context"
	"github.com/tiny-blob/tinyblob/pkg/controller"
	"github.com/tiny-blob/tinyblob/templates"
)

type errorHandler struct {
	controller.Controller
}

func (e *errorHandler) Get(err error, ctx echo.Context) {
	if ctx.Response().Committed || context.IsCanceledError(err) {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code >= 500 {
		ctx.Logger().Error(err)
	} else {
		ctx.Logger().Info(err)
	}

	page := controller.NewPage(ctx)
	page.Title = http.StatusText(code)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageError
	page.StatusCode = code
	page.HTMX.Request.Enabled = false

	if err = e.RenderPage(ctx, page); err != nil {
		ctx.Logger().Error(err)
	}
}
