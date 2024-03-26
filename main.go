package main

import (
	. "bus"
	. "dispatcher"
	. "events"
	. "examples"
	"fmt"
	"reflect"
	"sync"
)

func producer(w int, bus *Bus, p chan<- string, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	event := NewEvent(Projection{}, EventArgs{1: "1"})
	bus.Subscribe(event)
	fmt.Printf("Produced by worker %d \n", w)
	p <- fmt.Sprintf("%d", w)
}

func consumer(bus *Bus, p <-chan string, done chan bool) {
	for m := range p {
		fmt.Println("consumed:", m)
	}
	bus.Publish()
	done <- true
}

func main() {
	dispatcher := Dispatcher{
		reflect.TypeOf(Projection{}): Example,
	}
	bus := NewBus(&dispatcher)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	done := make(chan bool)
	producerQ := make(chan string)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go producer(i, bus, producerQ, &mutex, &wg)
	}
	go consumer(bus, producerQ, done)
	wg.Wait()

	close(producerQ)
	<-done
}
