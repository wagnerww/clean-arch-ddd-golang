package consumer

import eventStore "github.com/wagnerww/go-clean-arch-implement/domain/event-store"

type EventStoreConsumerInterface interface {
	Persist(repository eventStore.EventStoreRepositoryInterface)
}
