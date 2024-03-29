package store

import (
	. "events"
	. "stream"
	"sync"
)

func (store *Store) Publish(node chan Stream, wg *sync.WaitGroup, event Event, index int) chan Stream {
	store.Send(node)

	go Subscribe(node, event, wg, index)
	wg.Wait()
	return node
}

func Subscribe(nodeChan chan Stream, event Event, ws *sync.WaitGroup, j int) {
	data := NewStream(event, j)
	defer ws.Done()
	data.Append(nodeChan)
}
