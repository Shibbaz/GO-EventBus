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

func main() {
	arena := nuke.NewMonotonicArena(512*1024, 80)

	defer arena.Reset(true)
	var wg sync.WaitGroup
	start := time.Now()
	const SERVER_NUM = 100000
	store := Store{
		Dispatcher: &EventsDispatcher,
	}
	chanSlice := nuke.MakeSlice[chan Stream](arena, 0, SERVER_NUM)

	nodeRef := nuke.New[chan Stream](arena)
	*nodeRef = make(chan Stream)
	for i := 0; i < SERVER_NUM; i++ {
		wg.Add(1)
		event := NewEvent(EventArgs{"id": i, "price": 200000}, HouseWasSold{})
		*nodeRef = store.Publish(*nodeRef, &wg, event, i)
		chanSlice = nuke.SliceAppend(arena, chanSlice, *nodeRef)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	time.Sleep(200 * time.Millisecond)
}
