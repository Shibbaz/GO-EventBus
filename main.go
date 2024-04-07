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
				URLs: []string{"stun:stun.l.google.com:19302"},
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

func (eventstore *EventStore) Publish(channel string, event *Event) *webrtc.DataChannel {
	dc, err := eventstore.connection.CreateDataChannel(channel, nil)
	if err != nil {
		panic(err)
	}
	var result Event

	var buffer bytes.Buffer
	dec := gob.NewDecoder(&buffer)
	eventstore.connection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String()) //nolint
	})

	eventstore.connection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		dc.OnOpen(func() {
			log.Println("Event sourcing was initialized")
			enc := gob.NewEncoder(&buffer)
			err := enc.Encode(event)
			if err != nil {
				panic(err)
			}
			dc.Send(buffer.Bytes())
		})
		dc.OnMessage(func(msg webrtc.DataChannelMessage) {
			err := dec.Decode(&result)
			if err != nil {
				panic(err)
			}
			log.Printf("Event %d was published\n", result.Id)

		})
	})
	return dc
}

func main() {
	store := newEventStore()
	event := &Event{
		Id:         "1",
		Projection: "HouseWasSold",
		Args:       map[string]any{"price": 100},
	}
	dc := store.Publish("event", event)

	state := dc.ReadyState()
	fmt.Println(state)
	err := dc.Send([]byte{})
	if err != nil {
		panic(err)
	}
}
