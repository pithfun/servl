package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "{ok}")
}

func GetMeeting(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
