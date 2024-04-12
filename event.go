package GOEventBus

import (
	"reflect"

	"github.com/google/uuid"
)

type Event struct {
	Id         string
	Projection any
	Args       map[string]any
}

func NewEvent(projection any, args map[string]any) Event {
	id := uuid.New().String()
	return Event{
		Id:         id,
		Projection: reflect.TypeOf(projection).String(),
		Args:       args,
	}
}
