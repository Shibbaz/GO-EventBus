package bus

import (
	. "events"
	"fmt"
	"sync"
	"time"
)

func (bus *Bus) Work(event *Event, i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	projection := event.Projection
	fn := bus.Dispatcher.Get(projection)
	args := event.EventArgs
	status, _ := fn(args)
	bus.Channels[i] <- map[string]any{
		"id":     i,
		"status": status.Metadata,
	}
	event.Status = true

}

func (bus *Bus) Handle(size int) error {
	for i := 0; i < size; i++ {
		select {
		case event := <-bus.Channels[i]:
			fmt.Printf("Event %d succeded. Receiving data -> '%s'\n", event["id"], event["status"])
		default:
			return fmt.Errorf("no communication ready")
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func (bus *Bus) Publish() error {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	size := len(bus.Events)
	if len(bus.Events) == 0 {
		return nil
	}
	events := *NewEventList()
	for i := 0; i < size; i++ {
		if bus.Events[i].Status != true {
			bus.Channels[i] = make(chan map[string]any, 1)
			events = append(events, bus.Events[i])
			wg.Add(1)
		}
	}
	bus.Events = events
	size = len(bus.Events)
	for i := 0; i < size; i++ {
		go bus.Work(&events[i], i, &wg, &mutex)
	}
	wg.Wait()
	bus.Handle(size)
	return nil
}
