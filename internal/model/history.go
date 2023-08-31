package model

import "time"

type History struct {
	Id     int       `json:"id"`
	UserID string    `json:"user_id"`
	OpType string    `json:"operation_type"`
	OpDone time.Time `json:"omitempty"`
}
