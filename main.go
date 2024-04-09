package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/pion/webrtc/v2"
)

type Event struct {
	Id         string
	Projection any
	Args       map[string]any
}

func newEvent(projection any, args map[string]any) Event {
	id := uuid.New().String()
	return Event{
		Id:         id,
		Projection: reflect.TypeOf(projection).String(),
		Args:       args,
	}
}

type Dispatcher map[string]func(map[string]any)

type EventStore struct {
	dispatcher    Dispatcher
	OnDescription chan string
	OnBye         chan bool

	pc *webrtc.PeerConnection
}

func newEventStore(dispatcher Dispatcher) *EventStore {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{},
	}

	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	return &EventStore{
		dispatcher:    dispatcher,
		OnDescription: make(chan string),
		pc:            pc,
	}
}

func (eventstore *EventStore) Subscriber(event Event, channel string) {
	ordered := false
	maxRetransmits := uint16(0)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	options := &webrtc.DataChannelInit{
		Ordered:        &ordered,
		MaxRetransmits: &maxRetransmits,
	}
	dc, err := eventstore.pc.CreateDataChannel(channel, options)
	if err != nil {
		panic(err)
	}
	dc.OnOpen(func() {
		log.Printf("Channel %s was opened", dc.Label())
		err := enc.Encode(event)
		if err != nil {
			panic(err)
		}
		dc.Send(buffer.Bytes())
	})
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		log.Println("Successful")
	})
	offer, err := eventstore.pc.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	eventstore.pc.SetLocalDescription(offer)
	desc, err := json.Marshal(offer)
	if err != nil {
		panic(err)
	}

	go func() {
		eventstore.OnDescription <- string(desc)
	}()
}

func (eventstore *EventStore) Publish(event string) {
	var desc webrtc.SessionDescription
	dbyte := []byte(event)
	err := json.Unmarshal(dbyte, &desc)
	if err != nil {
		panic(err)
	}

	// Apply the desc as the remote description
	err = eventstore.pc.SetRemoteDescription(desc)
	if err != nil {
		panic(err)
	}

	// Set callback for new data channels
	eventstore.pc.OnDataChannel(func(dc *webrtc.DataChannel) {
		// Register channel opening handling
		dc.OnOpen(func() {})

		// Register the OnMessage to handle incoming messages
		dc.OnMessage(func(dcMsg webrtc.DataChannelMessage) {
			go func(msg webrtc.DataChannelMessage) {
				var result Event
				buffer := bytes.NewBuffer(msg.Data)
				dec := gob.NewDecoder(buffer)
				err := dec.Decode(&result)
				if err != nil {
					panic(err)
				}
				eventstore.dispatcher[result.Projection.(string)](result.Args)
				log.Printf("Event id of %s was published from channel '%s'", result.Id, dc.Label())
				dc.Send([]byte{})
			}(dcMsg)
		})

	})

	answer, err := eventstore.pc.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	eventstore.pc.SetLocalDescription(answer)
	desc2, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	go func() {
		eventstore.OnDescription <- string(desc2)
	}()
}

type HouseWasSold struct{}

func main() {
	dispatcher := Dispatcher{
		"main.HouseWasSold": func(m map[string]any) {
			fmt.Println(m)
		},
	}
	eventstore1 := newEventStore(dispatcher)
	eventstore2 := newEventStore(dispatcher)

	eventstore1.Subscriber(newEvent(
		HouseWasSold{},
		map[string]any{
			"key": 1,
		},
	), "eventstore")
	for {
		select {
		case event := <-eventstore1.OnDescription:
			eventstore2.Publish(event)
		case <-eventstore1.OnBye:
		case event := <-eventstore2.OnDescription:
			eventstore1.Publish(event)
		case <-eventstore1.OnBye:
		}

	}
}
