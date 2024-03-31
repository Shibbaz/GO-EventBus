package GOEventBus

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/ortuman/nuke"
)

// Store has its own arena, pointer to Event, batchSize
type Store struct {
	dispatcher *Dispatcher
	arena      nuke.Arena
	addr       uintptr
	events     *[]Event
	batchSize  int
}

// Store appends event to []Event
func (store *Store) Subscribe(event Event) Store {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	*store.events = append(*store.events, event)
	return *store
}

// Store recognizes events' action by projection type
func (store *Store) Publish(event *Event) error {
	t := reflect.TypeOf(event.projection)
	fn := (*(store.dispatcher))
	if fn == nil {
		return fmt.Errorf("cannot process an event")
	}
	if event.args == nil {
		return fmt.Errorf("cannot process an event")
	}
	fn[t](event.args)
	return nil
}

// Store broadcasts all events in the Store
func (store *Store) Broadcast() {
	defer store.arena.Reset(true)
	for _, event := range *store.events {
		if store == nil {
			continue
		}
		store.Publish(&event)
	}
}

func NewStore(dispatcher *Dispatcher, addr uintptr, bufferSize int, bufferCount int, batchSize int) *Store {
	arena := nuke.NewConcurrentArena(
		nuke.NewMonotonicArena(16*16, 20),
	)
	events := nuke.MakeSlice[Event](arena, 0, batchSize)

	return &Store{
		dispatcher: dispatcher,
		arena:      arena,
		addr:       addr,
		events:     &events,
		batchSize:  batchSize,
	}
}
