package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/pion/webrtc/v2"
)

type Event struct {
	Id         string
	Projection string
	Args       map[string]any
}

type EventStoreListener struct {
	OnEvent chan Event
	done    chan bool
}
type EventStore struct {
	connection *webrtc.PeerConnection
	listener   EventStoreListener
	dataChs    map[string]*webrtc.DataChannel
}

func newEventStore() *EventStore {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{
					"stun:127.0.0.1:3478",
				},
			},
		},
	}
	connection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	return &EventStore{
		connection: connection,
		listener: EventStoreListener{
			OnEvent: make(chan Event),
			done:    make(chan bool),
		},
		dataChs: map[string]*webrtc.DataChannel{},
	}
}

type HouseWasSold struct{}

func (eventstore *EventStore) Publish(channel string) *webrtc.DataChannel {
	dc, err := eventstore.connection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}
	var result Event

	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)
	eventstore.connection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String()) //nolint
	})

	dc.OnOpen(func() {
		log.Println("Event sourcing was initialized")
		if err != nil {
			panic(err)
		}
	})
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		err := dec.Decode(&result)
		if err != nil {
			panic(err)
		}
		log.Printf("Message ([]byte) from DataChannel '%s' with length %d\n", dc.Label())

	})
	return dc
}

func main() {
	store := newEventStore()

	dc := store.Publish("event")
	event := &Event{
		Id:         "1",
		Projection: "HouseWasSold",
		Args:       map[string]any{"price": 100},
	}
	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(event)
	if err != nil {
		panic(err)
	}
	err = dc.Send(buffer.Bytes())
	if err != nil {
		panic(err)
	}
}
