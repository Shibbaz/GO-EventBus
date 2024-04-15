package GOEventBus

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"log"

	"github.com/pion/webrtc/v2"
)

func (eventstore *EventStoreNode) Publish(event string, db *sql.DB) {
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

				data, err := eventstore.dispatcher[result.Projection.(string)](result.Args)
				if err != nil {
					ctx := context.Background()
					tx, err := db.BeginTx(ctx, nil)
					if err != nil {
						return
					}

					_, err = tx.ExecContext(ctx, "INSERT INTO events (event_id, projection, metadata) VALUES (?, ?, ?)", result.Id, result.Projection.(string), data)
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
