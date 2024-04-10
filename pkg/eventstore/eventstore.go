package eventstore

import (
	. "pkg"
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
}

type EventStore struct {
	Dispatcher *Dispatcher
	Done       chan bool
	events     sync.Pool
}

func NewEventStore(dispatcher *Dispatcher) *EventStore {
	return &EventStore{
		Dispatcher: dispatcher,
		events: sync.Pool{
			New: func() interface{} {
				return nil
			},
		},
	}
}

func (eventstore *EventStore) Publish(event Event) {
	eventstore.events.Put(event)
}

func (eventstore *EventStore) Broadcast() error {
	left := NewEventStoreNode(*eventstore.Dispatcher)
	right := NewEventStoreNode(*eventstore.Dispatcher)
	var done bool = false
	var mutex sync.Mutex
	var wg = sync.WaitGroup{}
	for {
		wg.Add(1)
		go func() {
			for done != true {
				go func() {
					mutex.Lock()
					curr := eventstore.events.Get()
					if curr == nil {
						done = true
						return
					}
					left.Subscribe(curr.(Event))
					mutex.Unlock()
				}()

			}
			wg.Done()
		}()
		wg.Wait()

		go func() {
			select {
			case event := <-left.Listner.OnDescription:
				right.Publish(event)
			case <-left.Listner.OnBye:
				eventstore.Done <- true
			case event := <-right.Listner.OnDescription:
				left.Publish(event)
			case <-right.Listner.OnBye:
			}
		}()
	}
}
