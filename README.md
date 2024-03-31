### GOEventBus
> Simple eventbus lib in Go, using Nuke memory arena to handle garbage collection for each batch of Events<br /><br />
To download module
```
go get github.com/Shibbaz/GOEventBus
```

Example
You can publish event just by
- create projection
	```
	type HousewasSold struct{}
	```
- create dispatcher
	```
	dispatcher := &bus.Dispatcher{
		reflect.TypeOf(bus.HouseWasSold{}): bus.Example,
	}
	```

- create and publish an event

	```
	store := bus.NewStore(dispatcher, 3, 16, 16, 3)
	event := bus.NewEvent(bus.HouseWasSold{}, bus.EventArgs{"price": 3})
	store.Publish(event)
	```

- Broadcast few events at one time
  
	```
	event1 := bus.NewEvent(bus.HouseWasSold{}, bus.EventArgs{"price": 1})
	event2 := bus.NewEvent(bus.HouseWasSold{}, bus.EventArgs{"price": 3})
	store.Subscribe(*event1)
	store.Subscribe(*event2)
	store.Broadcast()
	```

> bus.HouseWasSold{} is projection, bus.Example is func(args EventArgs) error


Look into examples.go to check out example Producer and Consumer case. 
> You can publish events like this where producer caches up every each users' action that produces event.<br />
> In this example, We produce the same event 100k times.
```
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
```

- Example of main.go program

```
package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	bus "github.com/Shibbaz/GOEventBus"

	"github.com/ortuman/nuke"
)

func main() {
	arena := nuke.NewMonotonicArena(32*32, 80)
	defer arena.Reset(true)

	storeRef := nuke.New[bus.Store](arena)
	now := time.Now()
	batches := make(chan bus.Store)
	dispatcher := &bus.Dispatcher{
		reflect.TypeOf(bus.HouseWasSold{}): bus.Example,
	}
	var wp sync.WaitGroup
	var wc sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wp.Add(1)
		*storeRef = *bus.NewStore(dispatcher, 3, 16, 16, 3)

		go bus.Producer(i, *storeRef, &wp, batches)
		wc.Add(1)
		go bus.Consumer(&wc, batches)
	}
	wp.Wait()
	close(batches)
	wc.Wait()

	elapsed := time.Since(now)
	fmt.Printf("elapsed %s", elapsed)
}
```
