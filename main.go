package main

import (
	. "events"
	. "examples"
	"reflect"
	. "store"
	"sync"
	"time"
	. "types"
)

type EventsStoreRepository struct {
	Store            *EventsStore
	EventsStoreMutex sync.Mutex
	EventsStoreWg    sync.WaitGroup
}

func NewEventsStoreRepository(dispatcher *EventDispatcher) *EventsStoreRepository {
	return &EventsStoreRepository{
		Store: &EventsStore{
			Dispatcher: *dispatcher,
		},
		EventsStoreMutex: sync.Mutex{},
		EventsStoreWg:    sync.WaitGroup{},
	}

}

func (repository *EventsStoreRepository) call() {
	repository.EventsStoreMutex.Lock()
	defer repository.EventsStoreWg.Done()
	defer repository.EventsStoreMutex.Unlock()

	var subscribersWg sync.WaitGroup
	var publisherWg sync.WaitGroup
	var mutex sync.Mutex

	subscribersWg.Add(1)
	subscriber := make(chan bool)
	// may be http request that executes below statement and go routine
	event := NewEvent(*NewProjection(), EventArgs{"1": 1, "2": 2})
	go repository.Store.SubscribeNewEvent(event, subscriber, &mutex, &subscribersWg)
	//
	subscribersWg.Wait()
	time.Sleep(100 * time.Millisecond)

	publisher := make(chan bool)
	publisherWg.Add(1)

	go repository.Store.PublishNewEvent(publisher, &mutex, &publisherWg)
	publisherWg.Wait()
}

func main() {
	dispatcher := NewEventDispatcher(
		map[reflect.Type]func(EventArgs) (Status, error){
			reflect.TypeOf(Projection{}): Example,
		},
	)
	repository := NewEventsStoreRepository(dispatcher)
	repository.EventsStoreWg.Add(1)
	go repository.call()
	repository.EventsStoreWg.Wait()
}
