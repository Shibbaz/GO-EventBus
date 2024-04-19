package GOEventBus

import (
	"encoding/json"
)

// Usecase of datachannel, event is encoded into bytes and send to another peer (look into Publish func)
func (eventstore *EventStoreNode) Subscribe(event Event) {

	eventstore.HandleDataChannel("subscribe", &event)
	offer, err := eventstore.connection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	eventstore.connection.SetLocalDescription(offer)
	desc, err := json.Marshal(offer)
	if err != nil {
		panic(err)
	}

	go func() {
		eventstore.Listner.OnDescription <- string(desc)
	}()
}
