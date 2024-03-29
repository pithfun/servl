package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/tiny-blob/tinyblob/pkg/controller"
	"github.com/tiny-blob/tinyblob/templates"
)

type (
	dashboard struct {
		controller.Controller
	}
)

func (c *dashboard) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageDashboard
	page.Title = "Dashboard"

	return c.RenderPage(ctx, page)
}
