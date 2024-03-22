package types

import "time"

type EventArgs map[string]any

type Projection struct {
	TimeStamp time.Time
}

func NewProjection() *Projection {
	return &Projection{
		TimeStamp: time.Now().UTC(),
	}
}

type Status struct {
	Code    int
	Message string
}

func NewStatus(code int, msg string) *Status {
	return &Status{
		Code:    code,
		Message: msg,
	}
}
