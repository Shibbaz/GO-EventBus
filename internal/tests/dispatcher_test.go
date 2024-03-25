package tests

import (
	. "dispatcher"
	. "events"
	. "examples"
	"reflect"
	"testing"
)

func TestDispatcherGet(t *testing.T) {

	dispatcher := Dispatcher{
		reflect.TypeOf(Projection{}): Example,
	}
	got, _ := dispatcher.Get(Projection{})(EventArgs{1: "1"})
	want, _ := Example(EventArgs{1: "1"})
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
