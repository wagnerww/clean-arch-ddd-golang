package events

type EventBusEventStoreInterface interface {
	Send(aggregate string, aggregateId string, action string, payload any)
}
