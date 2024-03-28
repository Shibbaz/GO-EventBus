package stream

import (
	. "events"
)

func NewStream(event Event, id int) Stream {
	channel := make(chan Stream)
	data := Stream{
		Id:    id,
		Nodes: make([]chan Stream, 0),
		Event: event,
	}
	go func(consumer chan Stream, d *Stream) {
		consumer <- *d
	}(channel, &data)
	data.Nodes = append(data.Nodes, channel)
	return data
}
