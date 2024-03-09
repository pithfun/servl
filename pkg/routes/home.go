package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/tiny-blob/tinyblob/pkg/controller"
	"github.com/tiny-blob/tinyblob/templates"
)

type (
	home struct {
		controller.Controller
	}
)

func (c *home) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Cache.Enabled = true
	page.Layout = templates.LayoutMain
	page.Name = templates.PageHome
	page.Metatags.Description = "A Player vs Player decentralized prediction market "
	page.Metatags.Keywords = []string{
		"Crypto",
		"PvP trading",
		"Decentralized",
		"Bitcoin",
		"PvP prediction market",
	}
	page.Title = "A PvP prediction marketplace"

	return c.RenderPage(ctx, page)
}
