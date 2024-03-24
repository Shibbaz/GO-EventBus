package store

import (
	. "events"
	"fmt"
	"reflect"
	"sync"
	"time"
)

func (store *EventsStore) AsyncPublish(channel chan string, event *Event, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	fn := store.GetFunc(event.Projection)
	typeOf := reflect.TypeOf(event.Projection).Name()
	fn(event.Args)
	event.Status = true
	channel <- typeOf
}

func (store *EventsStore) Publish() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	channels := make(map[int](chan string))
	if len(store.Stream) == 0 {
		wg.Wait()
	}
	size := len(store.Stream)
	for i := 0; i < size; i++ {
		if store.Stream[i].Status == false {
			channels[i] = make(chan string, 1)
			wg.Add(1)
		}
	}

	for i, event := range store.Stream {
		if event.Status == false {
			go store.AsyncPublish(channels[i], &event, &mutex, &wg)
		}
	}
	wg.Wait()
	store.Stream = []Event{}
	for i := 0; i < size; i++ {
		select {
		case value := <-channels[i]:
			fmt.Printf("Event type of %v succeded\n", value)
		default:
			fmt.Printf("No communication ready\n")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
