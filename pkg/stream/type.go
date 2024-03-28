package stream

import (
	. "events"
)

type Stream struct {
	Nodes []chan Stream
	Id    int
	Event Event
}
