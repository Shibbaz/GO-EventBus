package GOEventBus

import (
	"github.com/pion/webrtc/v2"
)

// EventStoreNode constructor
func NewEventStoreNode(dispatcher Dispatcher) *EventStoreNode {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{},
	}

	connection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	return &EventStoreNode{
		Listner: EventStoreListener{
			OnDescription: make(chan string),
			OnBye:         make(chan bool),
		},
		dispatcher: dispatcher,
		connection: connection,
	}
}
