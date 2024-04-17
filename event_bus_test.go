package GOEventBus

import (
	"fmt"
	"reflect"
	"testing"
)

type HouseWasSold struct{}

func TestNewEventStore(t *testing.T) {
	dispatcher := Dispatcher{
		"tests.HouseWasSold": func(m *map[string]any) (map[string]any, error) {
			fmt.Println(m)

			return *m, nil
		},
	}
	got := &EventStore{Dispatcher: &dispatcher}
	want := NewEventStore(&dispatcher)
	if reflect.DeepEqual(got, want) {
		t.Errorf("EventStore wants %v, got %v", got, want)
	}
}

func TestNewEvent(t *testing.T) {
	args := map[string]any{
		"price": 100,
	}
	event := NewEvent(HouseWasSold{}, args)
	if !reflect.DeepEqual(args, event.Args) {
		t.Errorf("Args are not correrct got %T, wanted %T", event.Args, args)

	}
}

func TestEventStorePublish(t *testing.T) {
	dispatcher := Dispatcher{
		"tests.HouseWasSold": func(m *map[string]any) (map[string]any, error) {
			fmt.Println(m)

			return *m, nil
		},
	}
	eventstore := NewEventStore(&dispatcher)
	args := map[string]any{
		"price": 100,
	}
	wanted := NewEvent(HouseWasSold{}, args)
	eventstore.Publish(wanted)
	got, valid := eventstore.GetEvent().(Event)
	if valid != true {
		t.Errorf("Event was not published, got %v, expected %v", got, wanted)

	}
}
