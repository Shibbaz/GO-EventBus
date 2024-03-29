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
	const SERVER_NUM = 100000
	store := Store{
		Dispatcher: &EventsDispatcher,
	}
	var node chan Stream
	node = make(chan Stream)
	for i := 0; i < SERVER_NUM; i++ {
		wg.Add(1)
		event := NewEvent(EventArgs{"id": i, "price": 200000}, HouseWasSold{})
		node = store.Publish(node, &wg, event, i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	time.Sleep(200 * time.Millisecond)

}
