package pkg

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Server struct {
	Conn       *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
	Dispatcher Dispatcher
}

func (server *Server) EventStoreSetup() {
	q, err := server.Channel.QueueDeclare(
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	FailOnError(err, "Cannot declare queue")
	server.Queue = q

}

func NewServer(dispatcher Dispatcher) *Server {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return &Server{
		Conn:       conn,
		Dispatcher: dispatcher,
		Channel:    ch,
	}
}
