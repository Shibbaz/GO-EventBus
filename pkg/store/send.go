package store

import (
	. "stream"
	"sync"
)

func (store *Store) Exec(index int, node *Stream, wg *sync.WaitGroup) {
	defer wg.Done()
	item := <-node.Nodes[index]
	store.Dispatcher.Get(item.Event.Projection)(item.Event.Args)
	go store.Send(node.Nodes[index])
}

func (store *Store) Send(data chan Stream) {
	go func(d chan Stream) {
		for node := range d {
			var wg sync.WaitGroup

			for i := 0; i < len(node.Nodes); i++ {
				wg.Add(1)
				go store.Exec(i, &node, &wg)
			}
			wg.Wait()

		}
	}(data)
}
