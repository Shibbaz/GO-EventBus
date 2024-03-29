package examples

import (
	. "events"
	"log"
	"net/http"
	. "store"
	. "stream"
	"sync"
)

type EventHandler struct {
	Store Store
	Node  chan Stream
}

func (handler *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := 0
	var wg sync.WaitGroup
	wg.Add(1)

	event := NewEvent(EventArgs{"id": i, "price": 200000}, HouseWasSold{})
	handler.Node = handler.Store.Publish(handler.Node, &wg, event, i)
	w.Write([]byte("The time is: "))
	wg.Wait()

}

func Server() *http.ServeMux {
	mux := http.NewServeMux()
	handler := EventHandler{
		Store: Store{
			Dispatcher: &EventsDispatcher,
		},
		Node: make(chan Stream),
	}

	mux.Handle("/event", &handler)

	log.Print("Listening...")
	return mux
}
