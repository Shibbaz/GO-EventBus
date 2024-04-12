### GOEventBus
> Simple eventbus lib in Go<br /><br />
- to download
```
get github.com/Shibbaz/GOEventBus
```

Examples
You can publish event just by
- create projection
```
type HouseSold struct{}
```
- create dispatcher
```
	dispatcher := Dispatcher{
		"main.HouseWasSold": func(m map[string]any) {
		},
	}
```
- create eventstore
```
  eventstore := NewEventStore(&dispatcher)
```
- create event
```
  event := NewEvent(
  			HouseWasSold{},
  			map[string]any{
  				"price": 1 * 100,
  			},
  		)
```
- publish it
```
  eventstore.Publish(event)
```


- run loop, you can put any eventsource within eventstore.Run func. For example listening to server and handler.
```
	eventstore.Run(func() {
		eventstore.Publish(NewEvent(
			HouseWasSold{},
			map[string]any{
				"price": 1 * 100,
			},
		))
	})
```



- Example of main.go program
```
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
```
