package store

import (
	. "events"
	. "stream"
	"sync"
)

func (store *Store) Publish(wg *sync.WaitGroup, event Event, index int) {
	node := make(chan Stream, 1)
	go Subscribe(node, event, wg, index)
	store.Send(node)
	wg.Wait()
	close(node)
}

func Subscribe(nodeChan chan Stream, event Event, ws *sync.WaitGroup, j int) {
	data := NewStream(event, j)
	defer ws.Done()
	data.Append(nodeChan)
}
