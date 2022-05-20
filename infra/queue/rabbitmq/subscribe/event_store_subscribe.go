package subscribe

import (
	"flag"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

type EventStoreRabbit struct {
	ch *amqp.Channel
}

func NewEventStoreRabbit(ch *amqp.Channel) *EventStoreRabbit {
	return &EventStoreRabbit{
		ch: ch,
	}
}

func (e *EventStoreRabbit) Persist() {

	consumerTag := flag.String("event-store", "simple-consumer", "AMQP consumer tag (should not be blank)")

	deliveries, err := e.ch.Consume(
		"events",
		*consumerTag,
		false, // noAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
	)

	if err != nil {
		log.Println("erro", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range deliveries {

			n, err := strconv.Atoi(string(d.Body))
			log.Println(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			//response := fib(n)

			/*err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			log.Println(err, "Failed to publish a message")*/

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever

}
