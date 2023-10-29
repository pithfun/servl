package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	Code     int         `json:"code"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"`
}

func GetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "{ok}")
}
