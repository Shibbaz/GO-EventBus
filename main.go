package main

import (
	. "events"
	. "examples"
	"fmt"
	"log"
	"net/http"
	. "store"
	. "stream"
	"sync"
	"time"

	"github.com/gorilla/mux"
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
	handler := EventHandler{
		Store: Store{
			Dispatcher: &EventsDispatcher,
		},
		Node: make(chan Stream),
	}
	mux := mux.NewRouter()

	mux.Handle("/event", &handler)

	srv := &http.Server{
		Handler: mux,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	log.Print("Listening...")

}
