package main

import (
	. "events"
	. "examples"
	"fmt"
	. "store"
	. "stream"
	"sync"
	"time"

	"github.com/ortuman/nuke"
)

func loop(i int, nodeRef chan Stream, wg *sync.WaitGroup, store *Store) {
	defer wg.Done()
	var wst sync.WaitGroup
	wst.Add(1)
	event := NewEvent(EventArgs{"id": i, "price": 200000}, HouseWasSold{})
	store.Publish(nodeRef, &wst, event, i)
}

func main() {
	arena := nuke.NewMonotonicArena(100*1024, 80)

	defer arena.Reset(true)
	var wg sync.WaitGroup
	start := time.Now()
	const SERVER_NUM = 100000
	store := Store{
		Dispatcher: &EventsDispatcher,
	}

	nodeRef := nuke.New[chan Stream](arena)
	*nodeRef = make(chan Stream)
	for i := 0; i < SERVER_NUM; i++ {
		wg.Add(1)
		go loop(i, *nodeRef, &wg, &store)
	}
	wg.Wait()
	close(*nodeRef)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
