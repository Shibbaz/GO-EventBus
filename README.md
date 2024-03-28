# Simple event bus library
> Developed in Go, implementing batching, dispatcher, using go routines
  - Projections describe what event does.
  - To declare projection
  ```
    type HouseWasSold struct{}
  ```
  - Creating events is simple
  ```
    event := NewEvent(HouseWasSold{}, args)
  ```
  - Create Event Dispatcher
  ```

  func HouseSell(EventArgs) error{
    fmt.Println("House was sold!")
    return nil
  }

  var EventsDispatcher = Dispatcher{
  	reflect.TypeOf(HouseWasSold{}): HouseSell,
  }
  ```
  - Use event store
  ```
    store := NewStore()
    store.Subscribe(event)
    store.Publish(dispatcher)
  ```
  - Use Batch Loader
  ```
    batch := Batch{}
    batchSize := 3
    batch = batch.Push(store.Events, batchSize)
    batch.Publish()
  ```
# Look into examples
```
var EventsDispatcher = Dispatcher{
	reflect.TypeOf(ExampleEvent{}): Example,
}

func Example(args EventArgs) error {
	fmt.Println(args)
	return nil
}

type ExampleEvent struct{}

func Subscribe(events chan []Event, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	store := NewStore()

	wp := sync.WaitGroup{}
	for j := 0; j < ProcessNum; j++ {
		wp.Add(1)
		go func() {
			defer wp.Done()
			event := NewEvent(ExampleEvent{}, EventArgs{1: 1})
			store.Subscribe(event)
		}()
		wp.Wait()
	}
	events <- store.Events
}

func Publish(event <-chan []Event, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	events := <-event
	batch := Batch{}
	batch.Publish(&events, BatchSize, &EventsDispatcher)
}

func main() {
	start := time.Now()

	events := make(Bus, ProcessNum)
	wp := sync.WaitGroup{}
	mutex := sync.Mutex{}

	wc := sync.WaitGroup{}
	cmutex := sync.Mutex{}

	for i := 0; i < ProcessNum; i++ {
		wp.Add(1)
		go Subscribe(events, &wp, &mutex)
		wc.Add(1)
		go Publish(events, &wc, &cmutex)
	}
	wp.Wait()
	wc.Wait()
	close(events)

	elapsed := time.Since(start)
	log.Printf("1000000 events took %s", elapsed)
}
```
