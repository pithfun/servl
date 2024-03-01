package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/tiny-blob/tinyblob/pkg/controller"
)

type (
	home struct {
		controller.Controller
	}
)

func (c *home) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Cache.Enabled = true
	page.Layout = "main"
	page.Name = "home"
	page.Metatags.Description = "Instantly execute complex crypto strategies at lightning speed."
	page.Metatags.Keywords = []string{
		"Crypto trading",
		"Solana trading",
		"Automated trading",
		"Lightning-fast transactions",
		"Seamless execution",
		"Crypto strategies",
		"Drag and drop builder",
		"Signal trading",
		"Trade executor",
	}
	page.Title = "Crypto tools for professionals"

	return c.RenderPage(ctx, page)
}
