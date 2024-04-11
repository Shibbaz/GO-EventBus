package main

import (
	. "eventstore"
	. "pkg"
)

type HouseWasSold struct{}

func main() {
	dispatcher := Dispatcher{
		"main.HouseWasSold": func(m map[string]any) {
		},
	}
	eventstore := NewEventStore(&dispatcher)
	eventstore.Run(func() {
		eventstore.Publish(NewEvent(
			HouseWasSold{},
			map[string]any{
				"price": 1 * 100,
			},
		))
	})
}
