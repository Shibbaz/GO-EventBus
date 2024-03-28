package main

import (
	. "events"
	. "examples"
	"fmt"
	. "store"
	. "stream"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	const SERVER_NUM = 200000
	store := Store{
		Dispatcher: &EventsDispatcher,
	}
	for i := 0; i < SERVER_NUM; i++ {
		wg.Add(1)
		node := make(chan Stream, 1)
		go func(nodeChan chan Stream, ws *sync.WaitGroup, j int) {
			event := NewEvent(EventArgs{"id": j, "price": 200000}, HouseWasSold{})
			data := NewStream(event, j, j)
			defer ws.Done()
			nodeChan <- data
			data.Append(nodeChan)
		}(node, &wg, i)

		wg.Wait()
		store.Send(node)

	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s", elapsed)
}
