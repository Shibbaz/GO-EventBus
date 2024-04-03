package main

import (
	"log"

	. "pkg"
)

func main() {
	server := NewServer(ExampleDispatcher)
	defer server.Conn.Close()

	defer server.Channel.Close()
	err := server.Channel.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	FailOnError(err, "Failed to declare exchange")

	server.EventStoreSetup()

	err = server.Channel.QueueBind(
		server.Queue.Name, // queue name
		"",                // routing key
		"logs",            // exchange
		false,
		nil,
	)
	FailOnError(err, "Failed to bind a queue")
	msgs, err := server.Channel.Consume(
		server.Queue.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for event := range msgs {
			server.Apply(&event)
		}
	}()

	log.Printf(" [*] Waiting for events. To exit press CTRL+C")
	<-forever
}
