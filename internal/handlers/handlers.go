package handlers

import (
	"net/http"
	"strconv"
	"time"

	"meetpanel/pkg/zoom"

	"github.com/google/uuid"
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

func GetMeeting(c echo.Context) error {
	ID := c.Param("id")

	meetingID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Invalid meeting ID",
		})
	}

	m := zoom.Meeting{
		AssistantID: "kFFvsJc-Q1OSxaJQLvaa_A",
		Agenda:      "Meeting with John Fetterman",
		CreatedAt:   time.Now().UTC(),
		ID:          meetingID,
		Topic:       "Scheduled fireside chat with John Fetterman",
		Type:        2,
		UUID:        uuid.New().String(),
	}

	return c.JSON(http.StatusOK, m)
}
