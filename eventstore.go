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
	Dispatcher *Dispatcher
	Done       chan bool
	events     sync.Pool
	Wg         sync.WaitGroup
	left       EventStoreNode
	right      EventStoreNode
	mutex      sync.Mutex
}

var EventStoreDB *sql.DB

func SetEventStoreDB(psqlInfo string) {
	var err error
	EventStoreDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func NewEventStore(dispatcher *Dispatcher) *EventStore {
	left := NewEventStoreNode(*dispatcher)
	right := NewEventStoreNode(*dispatcher)
	return &EventStore{
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
func (eventstore *EventStore) Query(projection string) map[string](map[string]any) {
	rows, err := EventStoreDB.Query("SELECT event_id, metadata FROM events where projection = %s;", projection)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	data := map[string](map[string]any){}
	for rows.Next() {
		var metadata []byte
		var event_id string
		if err := rows.Scan(&metadata); err != nil {
			log.Fatal(err)
		}
		if err := rows.Scan(&event_id); err != nil {
			log.Fatal(err)
		}
		data[event_id] = NewSerializer().Deserialize(metadata)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
func (eventstore *EventStore) Setup(dbname string) {
	_, err := EventStoreDB.Exec("create database " + dbname)
	if err != nil {
		log.Fatal(err)
	}

	_, err = EventStoreDB.Exec("CREATE TABLE IF NOT EXISTS events(event_id text primary key, projection text, metadata bytea)")

	if err != nil {
		log.Fatal(err)
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
		eventstore.right.Publish(event)
		event2 := <-eventstore.right.Listner.OnDescription
		eventstore.left.Publish(event2)
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
