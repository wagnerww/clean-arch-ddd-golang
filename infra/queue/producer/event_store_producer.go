package event_store

type EventStoreProducerInterface interface {
	Send(aggregate string, aggregateId string, action string, payload any)
}
