package model

import "time"

type User struct {
	ID          int
	Expire_date time.Time
}
