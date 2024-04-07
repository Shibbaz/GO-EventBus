package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"reflect"

	"github.com/pion/webrtc/v2"
)

type Event struct {
	id         string
	projection reflect.Type
	args       map[string]any
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

func (eventstore *EventStore) Publish(event *Event, channel string) {
	ordered := false
	maxRetransmits := uint16(0)

	options := &webrtc.DataChannelInit{
		Ordered:        &ordered,
		MaxRetransmits: &maxRetransmits,
	}

	dc, err := eventstore.connection.CreateDataChannel(channel, options)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	dec := gob.NewDecoder(&buffer)
	eventstore.connection.OnDataChannel(func(d *webrtc.DataChannel) {

		dc.OnOpen(func() {
			log.Println("Event of type HouseWasSold was sent")
			err := enc.Encode(event)
			if err != nil {
				log.Fatal("encode error:", err)
			}
			err = dc.Send(buffer.Bytes())
			if err != nil {
				panic(err)
			}
		})
		var ev Event
		dc.OnMessage(func(msg webrtc.DataChannelMessage) {
			err := dec.Decode(&ev)
			if err != nil {
				panic(err)
			}
			log.Printf("Message ([]byte) from DataChannel '%s' with length %d\n", dc.Label())

		})
	})

	fmt.Println(dc.ReadyState())

	// Send a message over the data channel
	message := []byte("Hello, world!")
	if err := dc.Send(message); err != nil {
		log.Fatal(err)
	}

	// Close the connection when the program is finished
	defer eventstore.connection.Close()

	eventstore.dataChs[dc.Label()] = dc

	offer, err := eventstore.connection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	eventstore.connection.SetLocalDescription(offer)

	go func() {
		eventstore.listener.OnEvent <- *event
	}()
}

func main() {
	store := newEventStore()

	store.Publish(&Event{
		id:         "1",
		projection: reflect.TypeOf(HouseWasSold{}),
		args:       map[string]any{"price": 100},
	}, "hello")
	for {
		// Block forever
		select {
		case event := <-store.listener.OnEvent:
			fmt.Printf("event id %s of type %s was published", event.id, event.projection.Name())
		case <-store.listener.done:
		}
	}
}
