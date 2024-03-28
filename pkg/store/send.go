package store

import (
	. "stream"
	"sync"
)

func (store *Store) Send(data chan Stream) {
	go func() {
		for node := range data {
			var wg sync.WaitGroup

			for i := 0; i < len(node.Nodes); i++ {
				wg.Add(1)

				go func(ws *sync.WaitGroup) {
					defer ws.Done()
					item := <-node.Nodes[i]
					store.Dispatcher.Get(item.Event.Projection)(item.Event.Args)
					store.Send(node.Nodes[i])
				}(&wg)
			}
			wg.Wait()

		}
	}()
}
