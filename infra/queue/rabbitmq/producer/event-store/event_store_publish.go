package publish

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	eventStore "github.com/wagnerww/go-clean-arch-implement/infra/event-store"
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

	event := eventStore.EventStore{
		ID:          uuid.New().String(),
		Aggregate:   aggregate,
		AggregateId: aggregateId,
		Action:      action,
		Payload:     payload,
	}

	body, _ := json.Marshal(event)

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
