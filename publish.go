package GOEventBus

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/pion/webrtc/v2"
)

func (eventstore *EventStoreNode) HandleDataChannel(id string, event *Event) {
	switch id {
	case "subscribe":
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
			dc.Send([]byte{})
			dc.Close()

		})
	case "publish":
		eventstore.connection.OnDataChannel(func(dc *webrtc.DataChannel) {
			dc.OnOpen(func() {
				dc.Send([]byte{})
			})

			dc.OnMessage(func(dcMsg webrtc.DataChannelMessage) {
				go func(msg webrtc.DataChannelMessage) {
					var result Event
					buffer := bytes.NewBuffer(msg.Data)
					dec := gob.NewDecoder(buffer)
					err := dec.Decode(&result)
					if err != nil {
						panic(err)
					}
					data, err := eventstore.dispatcher[result.Projection.(string)](&result.Args)
					if err != nil {
						panic(err)
					}
					serialized := NewSerializer().Serialize(data)

					if serialized != nil {

						ctx := context.Background()
						tx, err := EventStoreDB.BeginTx(ctx, nil)
						if err != nil {
							return
						}

						_, err = EventStoreDB.ExecContext(ctx, "INSERT INTO events (event_id, projection, metadata) VALUES ($1, $2, $3)", result.Id, result.Projection.(string), serialized)

						if err != nil {
							return
						}
						if err = tx.Commit(); err != nil {
							return
						}
						log.Printf("Event id of %s was published from channel '%s'", result.Id, dc.Label())
						dc.Send([]byte{})
					}

				}(dcMsg)
			})

		})
	}
}

// Sending bytes over datachannels to publish events and send it in database
// Serializing data before inserting it to database
func (eventstore *EventStoreNode) Publish(event string) {
	var desc webrtc.SessionDescription
	dbyte := []byte(event)
	err := json.Unmarshal(dbyte, &desc)
	if err != nil {
		panic(err)
	}

	err = eventstore.connection.SetRemoteDescription(desc)
	if err != nil {
		panic(err)
	}

	eventstore.HandleDataChannel("publish", nil)

	answer, err := eventstore.connection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	eventstore.connection.SetLocalDescription(answer)
	desc2, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	go func() {
		eventstore.Listner.OnDescription <- string(desc2)
	}()
}
