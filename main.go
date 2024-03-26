package main

import (
	. "bus"
	. "dispatcher"
	. "events"
	. "examples"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

func producer(w int, bus *Bus, p chan<- string, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	event := NewEvent(Projection{}, EventArgs{1: w})
	bus.Subscribe(event)
	fmt.Printf("Produced by worker %d \n", w)
	p <- fmt.Sprintf("%d", w)
}

func consumer(bus *Bus, p <-chan string, done chan bool) {
	for _ = range p {
		bus.Publish()
	}
	done <- true
}

func main() {
	start := time.Now()
	dispatcher := Dispatcher{
		reflect.TypeOf(Projection{}): Example,
	}
	bus := NewBus(&dispatcher)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	done := make(chan bool)
	producerQ := make(chan string)
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go producer(i, bus, producerQ, &mutex, &wg)
	}
	go consumer(bus, producerQ, done)
	wg.Wait()

	close(producerQ)
	<-done
	elapsed := time.Since(start)
	log.Printf("100 000 events took %s", elapsed)
}
