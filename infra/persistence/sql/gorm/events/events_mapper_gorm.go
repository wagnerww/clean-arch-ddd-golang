package orm

import (
	"encoding/json"

	"github.com/google/uuid"
	event_store "github.com/wagnerww/go-clean-arch-implement/domain/event-store"
)

func DomainToModel(domain event_store.EventStore) (eventStore EventStore) {

	payload, _ := json.Marshal(domain.Payload)

	model := EventStore{
		ID:           uuid.MustParse(domain.ID),
		Aggregate:    domain.Aggregate,
		AggregateId:  domain.AggregateId,
		Action:       domain.Action,
		IsSyncNeeded: domain.IsSyncNeeded,
		Payload:      string(payload),
	}

	return model

}
