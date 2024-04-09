package main

import (
	. "eventstore"
	. "pkg"
	"sync"
)

type HouseWasSold struct{}

func main() {
	dispatcher := Dispatcher{
		"main.HouseWasSold": func(m map[string]any) {
		},
	}
	eventstore := NewEventStore(&dispatcher)
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			go eventstore.Publish(NewEvent(
				HouseWasSold{},
				map[string]any{
					"price": i * 100,
				},
			))

		}
	}()
	wg.Wait()

	go eventstore.Broadcast()

	<-eventstore.Done

	close(eventstore.Done)

}
