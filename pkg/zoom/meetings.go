package zoom

import (
	"time"
)

type Meeting struct {
	AssistantID string    `json:"assistant_id,omitempty"` // The ID of the user who scheduled this meeting on behalf of the host.
	Agenda      string    `json:"agenda,omitempty"`       // Meeting description.
	CreatedAt   time.Time `json:"created_at"`             // Meeting creation time.
	ID          int64     `json:"id"`                     // Unique identifier of the meeting in long format(represented as int64 data type in JSON).
	Topic       string    `json:"topic,omitempty"`        // Meeting topic.
	Type        int       `json:"type"`                   // Meeting type.
	UUID        string    `json:"uuid"`                   // Unique identifier of the meeting. Each meeting instance will generate its own meeting UUID.
}
