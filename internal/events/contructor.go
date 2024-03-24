package events

import (
	. "types"
)

func NewEvent(projection any, args EventArgs) *Event {
	return &Event{
		Projection: projection,
		Args:       args,
		Status:     false,
	}
}
