package GOEventBus

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/pion/webrtc/v2"
)

type EventStoreListener struct {
	OnDescription chan string
	OnBye         chan bool
}
type EventStoreNode struct {
	connection *webrtc.PeerConnection
	dispatcher Dispatcher
	Listner    EventStoreListener
	DC         webrtc.DataChannel
}

type EventStore struct {
	DB         *sql.DB
	Dispatcher *Dispatcher
	Done       chan bool
	events     sync.Pool
	Wg         sync.WaitGroup
	left       EventStoreNode
	right      EventStoreNode
	mutex      sync.Mutex
}

func NewEventStore(dispatcher *Dispatcher, db *sql.DB) *EventStore {
	left := NewEventStoreNode(*dispatcher)
	right := NewEventStoreNode(*dispatcher)
	return &EventStore{
		DB:         db,
		left:       *left,
		right:      *right,
		Dispatcher: dispatcher,
		mutex:      sync.Mutex{},
		events: sync.Pool{
			New: func() interface{} {
				return nil
			},
		},
	}
}

func (eventstore *EventStore) GetEvent() any {
	return eventstore.events.Get()
}
func (eventstore *EventStore) Publish(event Event) {
	eventstore.events.Put(event)
}

func (eventstore *EventStore) Broadcast() error {
	eventstore.mutex.Lock()
	defer eventstore.mutex.Unlock()
	for {
		curr := eventstore.events.Get()
		if curr == nil {
			return fmt.Errorf("waiting for new events...")
		}
		eventstore.left.Subscribe(curr.(Event))
		event := <-eventstore.left.Listner.OnDescription
		eventstore.right.Publish(event, eventstore.DB)
		event2 := <-eventstore.right.Listner.OnDescription
		eventstore.left.Publish(event2, eventstore.DB)
	}
}

func (eventstore *EventStore) Run(EventsSource func()) {
	log.Println("EventStore initialized!")
	var mutex = sync.Mutex{}
datasource:
	{
		EventsSource()
	}
	go func() {
		mutex.Lock()
		err := eventstore.Broadcast()
		if err != nil {
			return
		}
		mutex.Unlock()
	}()
	goto datasource

}
