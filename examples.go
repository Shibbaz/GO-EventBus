package GOEventBus

import (
	"fmt"
	"sync"
)

type HouseWasSold struct{}

func Example(args EventArgs) error {
	fmt.Printf("House was sold for %d\n", args["price"])
	return nil
}

func Producer(i int, store Store, wg *sync.WaitGroup, channel chan Store) {
	defer wg.Done()
	event := NewEvent(HouseWasSold{}, EventArgs{"price": i})

	if len(*store.events)+1 == store.batchSize {
		store = *NewStore(store.dispatcher, store.addr, 16, 16, store.batchSize)
	}

	store.Subscribe(*event)

	channel <- store

}

func Consumer(wg *sync.WaitGroup, storeChan chan Store) {
	defer wg.Done()
	go func(channel chan Store) {
		for store := range channel {
			store.Broadcast()
		}
	}(storeChan)
}
