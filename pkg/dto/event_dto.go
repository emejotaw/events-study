package dto

import "time"

type Event struct {
	Name      string
	Code      int
	Message   string
	CreatedAt time.Time
}
