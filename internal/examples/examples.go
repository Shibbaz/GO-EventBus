package examples

import (
	. "batch"
	. "events"
	"fmt"
	. "helpers"
	"reflect"
	"sync"
)

var EventsDispatcher = Dispatcher{
	reflect.TypeOf(ExampleEvent{}): Example,
}

func Example(args EventArgs) error {
	fmt.Println(args)
	return nil
}

type ExampleEvent struct{}

func Subscribe(events chan []Event, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	batchevents := []Event{}
	wp := sync.WaitGroup{}
	for j := 0; j < ProcessNum; j++ {
		wp.Add(1)
		go func() {
			defer wp.Done()
			event := NewEvent(ExampleEvent{}, EventArgs{1: 1})
			batchevents = append(batchevents, event)
		}()
		wp.Wait()
	}
	events <- batchevents
}

func Publish(event <-chan []Event, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	events := <-event
	batch := Batch{}
	batch.Publish(&events, BatchSize, &EventsDispatcher)
}
