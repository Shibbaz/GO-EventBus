package bus

import (
	"fmt"
	"sync"
	"time"
)

/*
	type Bus struct {
		Events      []Event
		Subscribers []chan Subscriber
		Dispatcher  Dispatcher
	}
*/

func (bus *Bus) Publish(i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	projection := bus.Events[i].Projection
	fn := bus.Dispatcher.Get(projection)
	args := bus.Events[i].EventArgs
	status, _ := fn(args)
	bus.Channels[i] <- map[string]any{
		"id":     i,
		"status": status.Metadata,
	}

}

func (bus *Bus) Compose() {

	var wg sync.WaitGroup
	var mutex sync.Mutex
	size := len(bus.Events)
	if len(bus.Events) == 0 {
		wg.Wait()
	}
	for i := 0; i < size; i++ {
		bus.Channels[i] = make(chan map[string]any, 1)
		wg.Add(1)
	}
	for i := 0; i < size; i++ {
		go bus.Publish(i, &wg, &mutex)
	}
	wg.Wait()

	for i := 0; i < size; i++ {
		select {
		case event := <-bus.Channels[i]:
			fmt.Printf("Event %d succeded. Receiving data -> '%s'\n", event["id"], event["status"])
		default:
			fmt.Printf("No communication ready\n")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
