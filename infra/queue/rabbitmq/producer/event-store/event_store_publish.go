package publish

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	eventStore "github.com/wagnerww/go-clean-arch-implement/domain/event-store"
	event_store "github.com/wagnerww/go-clean-arch-implement/domain/event-store"
)

type EventStoreRabbit struct {
	ch     *amqp.Channel
	repoDb event_store.EventStoreRepositoryInterface
}

func NewEventStoreRabbit(ch *amqp.Channel, repoDb event_store.EventStoreRepositoryInterface) *EventStoreRabbit {
	return &EventStoreRabbit{
		ch:     ch,
		repoDb: repoDb,
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

	var err error
	if e.ch == nil {
		err = errors.New("Falha em conexão com RabbitMQ")
	}

	if err == nil {
		err = e.ch.Publish(
			"events_direct",
			"",
			false,
			false, amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			})
	}

	if err != nil {
		log.Println("erro_rabbit", err)
		e.repoDb.Create(event)
	}

}
