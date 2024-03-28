package examples

import (
	. "github.com/Shibbaz/GO-EventBus/pkg/batch"

	. "github.com/Shibbaz/GO-EventBus/pkg/events"

	"fmt"

	. "github.com/Shibbaz/GO-EventBus/pkg/helpers"

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

func EventsProducer() []Event {
	store := NewStore()

	wp := sync.WaitGroup{}
	for job := 0; job < ProcessNum; job++ {
		wp.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			event := NewEvent(ExampleEvent{}, EventArgs{1: 1})
			store.Subscribe(event)
		}(&wp)
		wp.Wait()
	}
	return store.Events
}

func Subscribe(eventsChannels chan []Event, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	var ws sync.WaitGroup
	ws.Add(1)
	go func(ws *sync.WaitGroup) {
		defer ws.Done()
		eventsChannels <- EventsProducer()
	}(&ws)
	ws.Wait()
}

func Publish(event chan []Event, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	batch := Batch{}
	batch.Publish(event, BatchSize, &EventsDispatcher)
}
