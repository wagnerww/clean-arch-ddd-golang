package event_store

type EventStoreRepositoryInterface interface {
	Create(entity EventStore) error
}
