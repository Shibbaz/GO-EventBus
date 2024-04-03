package pkg

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Agregate map[string]interface{}

type Event struct {
	Id   string
	Name string
	Args map[string]any
}

func (server *Server) Apply(event *amqp091.Delivery) {
	msg, err := Deserialize(event.Body)
	FailOnError(err, "Cannot deserialize")
	server.Dispatcher[msg["type"].(string)](msg["args"].(map[string]any))
	log.Printf("Received an event : %s", msg)
}

func (server *Server) Publish(event *Event) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msg, err := Serialize(Agregate{"id": event.Id, "type": event.Name, "args": event.Args})
	FailOnError(err, "Cannot serialize Event")
	err = server.Channel.PublishWithContext(ctx,
		"",                // exchange
		server.Queue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent Event %s\n", msg)
}
