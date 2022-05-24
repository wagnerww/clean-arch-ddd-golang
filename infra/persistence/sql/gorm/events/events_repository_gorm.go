package orm

import (
	event_store "github.com/wagnerww/go-clean-arch-implement/domain/event-store"
	"gorm.io/gorm"
)

type EventStoreRepositoryGorm struct {
	db *gorm.DB
}

func NewEventStoreRepositoryGorm(db *gorm.DB) *EventStoreRepositoryGorm {
	return &EventStoreRepositoryGorm{
		db: db,
	}
}

func (c *EventStoreRepositoryGorm) Create(entity event_store.EventStore) error {
	model := DomainToModel(entity)
	c.db.Create(&model)

	return nil
}
