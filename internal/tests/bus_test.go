package tests

import (
	. "bus"
	. "dispatcher"
	. "events"
	. "examples"
	"reflect"
	"testing"
)

func TestPublishBus(t *testing.T) {
	dispatcher := Dispatcher{
		reflect.TypeOf(Projection{}): Example,
	}
	bus := NewBus(&dispatcher)

	event := NewEvent(Projection{}, EventArgs{1: "1"})
	bus.Subscribe(event)
	got := bus.Compose()
	if !reflect.DeepEqual(nil, got) {
		t.Fatalf("wanted %v, got %v", nil, got)
	}
}

func TestSubscribeBus(t *testing.T) {

	type HouseWasBought struct{}
	type HouseWasSold struct{}
	dispatcher := Dispatcher{
		reflect.TypeOf(HouseWasBought{}): Example,
		reflect.TypeOf(HouseWasSold{}):   Example,
	}
	bus := NewBus(&dispatcher)

	event := NewEvent(HouseWasBought{}, EventArgs{1: "1"})
	bus.Subscribe(event)
	event = NewEvent(HouseWasSold{}, EventArgs{2: "2"})
	bus.Subscribe(event)

	size := len(bus.Events)
	got := 2
	if !reflect.DeepEqual(size, got) {
		t.Fatalf("wanted %v, got %v", nil, got)
	}
}
