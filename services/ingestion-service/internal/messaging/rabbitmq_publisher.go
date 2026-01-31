package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQPublisher(url string, queueName string) *RabbitMQPublisher {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("RabbitMQ bağlantı hatası: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Kanal açma hatası: %s", err)
	}

	_, err = ch.QueueDeclare(
		queueName, true, false, false, false, nil,
	)

	return &RabbitMQPublisher{conn: conn, channel: ch, queue: queueName}
}

func (p *RabbitMQPublisher) Publish(message []byte) error {
	return p.channel.Publish(
		"", p.queue, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}