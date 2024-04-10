package eventstore

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/pion/webrtc/v2"
)

func (eventstore *EventStoreNode) Subscribe(event Event) {
	ordered := false
	maxRetransmits := uint16(0)
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	options := &webrtc.DataChannelInit{
		Ordered:        &ordered,
		MaxRetransmits: &maxRetransmits,
	}
	dc, err := eventstore.connection.CreateDataChannel(uuid.New().String(), options)
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
		eventstore.Listner.OnBye <- true
	})
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
