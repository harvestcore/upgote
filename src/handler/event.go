package handler

import (
	"github.com/google/uuid"
)

type Event struct {
	Id     uuid.UUID
	From   uuid.UUID
	Type   uint8
	Action string
	Data   map[interface{}]interface{}
}
