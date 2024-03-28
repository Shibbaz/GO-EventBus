package main

import (
	"log"
	"sync"
	"time"

	. "github.com/Shibbaz/GO-EventBus/pkg/helpers"

	. "github.com/Shibbaz/GO-EventBus/pkg/examples"
)

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
