package model

import "time"

type UserSegments struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	SegmentNames []string  `json:"add_segment_names,omitempty"`
	ExpireDate   time.Time `json:"expire_date,omitempty"`
}
