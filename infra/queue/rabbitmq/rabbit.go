package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Channel, *amqp.Connection) {
	amqpServerURL := "amqp://guest:guest@localhost:5672/"
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	//defer connectRabbitMQ.Close()

	err = channelRabbitMQ.ExchangeDeclare(
		"events_direct", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	_, err = channelRabbitMQ.QueueDeclare(
		"events", // queue name
		false,    // durable
		false,    // auto delete
		false,    // exclusive
		false,    // no wait
		nil,      // arguments
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Binding queue %s to exchange %s with routing key %s", "events", "events_direct", "")

	channelRabbitMQ.QueueBind(
		"events",        // queue name
		"",              // routing key
		"events_direct", // exchange
		false,
		nil)

	//defer channelRabbitMQ.Close()

	return channelRabbitMQ, connectRabbitMQ

}
