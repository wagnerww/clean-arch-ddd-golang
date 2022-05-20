package orm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerModel struct {
	gorm.Model
}

type Customer struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"column:nome"`
	Email     string    `gorm:"column:email"`
	Active    bool      `gorm:"column:ativo"`
	Telephone int       `gorm:"column:telefone"`
	City      string    `gorm:"column:cidade"`
	Street    string    `gorm:"column:logradouro"`
	Number    int       `gorm:"column:numero"`
	CreatedAt time.Time `gorm:"column:criado_em"`
	UpdatedAt time.Time `gorm:"column:editado_em"`
	DeletedAt time.Time `gorm:"column:apagado_em"`
}

type Table interface {
	TableName() string
}

func (Customer) TableName() string {
	return "clientes"
}
