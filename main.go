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

func Setup(wg sync.WaitGroup, store *Store, index int) {
	wg.Add(1)
	node := make(chan Stream, 1)
	go func(nodeChan chan Stream, ws *sync.WaitGroup, j int) {
		event := NewEvent(EventArgs{"id": j, "price": 200000}, HouseWasSold{})
		data := NewStream(event, j)
		defer ws.Done()
		data.Append(nodeChan)
	}(node, &wg, index)
	store.Send(node)
	wg.Wait()
	close(node)
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	const SERVER_NUM = 100000
	store := Store{
		Dispatcher: &EventsDispatcher,
	}
	for i := 0; i < SERVER_NUM; i++ {
		go Setup(wg, &store, i)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	time.Sleep(200 * time.Millisecond)

}
