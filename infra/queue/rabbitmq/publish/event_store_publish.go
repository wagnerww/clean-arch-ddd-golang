package publish

import (
	"encoding/json"
	"log"

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

func (e *EventStoreRabbit) Send(aggregate string, aggregateId string, action string, payload any) {

	body, _ := json.Marshal(payload)

	err := e.ch.Publish(
		"events_direct",
		"",
		false,
		false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	log.Println("erro_rabbit", err)

}
