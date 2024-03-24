package store

import (
	. "events"
	"sync"
)

func (eventStore *EventsStore) SubscribeNewEvent(event *Event, subscriber chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	eventStore.Subscribe(event)
}

func (eventStore *EventsStore) PublishNewEvent(publisher chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	eventStore.Publish()
}
