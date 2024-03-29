package main

import (
	. "events"
	. "examples"
	"fmt"
	. "store"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	const SERVER_NUM = 1000
	store := Store{
		Dispatcher: &EventsDispatcher,
	}
	for i := 0; i < SERVER_NUM; i++ {
		wg.Add(1)
		event := NewEvent(EventArgs{"id": i, "price": 200000}, HouseWasSold{})
		store.Publish(&wg, event, i)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	time.Sleep(200 * time.Millisecond)

}
