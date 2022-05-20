package orm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventStoreModel struct {
	gorm.Model
}

type EventStore struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid"`
	Aggregate    string    `gorm:"column:agregado"`
	AggregateId  string    `gorm:"column:agregado_id"`
	Action       string    `gorm:"column:acao"`
	Payload      string    `gorm:"column:payload"`
	IsSyncNedeed bool      `gorm:"column:necessita_sincronizacao"`
	CreatedAt    time.Time `gorm:"column:criado_em"`
	UpdatedAt    time.Time `gorm:"column:editado_em"`
	DeletedAt    time.Time `gorm:"column:apagado_em"`
}

type Table interface {
	TableName() string
}

func (EventStore) TableName() string {
	return "eventos"
}
