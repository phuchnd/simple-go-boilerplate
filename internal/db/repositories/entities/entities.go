package entities

import (
	"strconv"
	"time"
)

type ID int64

// IDSetter abstracts an object that can set ID.
type IDSetter interface {
	SetID(id ID)
}

func (id ID) ToString() string {
	return strconv.FormatInt(int64(id), 10)
}

// Model is the base model.
type Model struct {
	ID        ID        `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func (e *Model) GetID() ID {
	return e.ID
}

func (e *Model) SetID(id ID) {
	e.ID = id
}

func (e *Model) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *Model) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}
