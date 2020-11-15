package event

import (
	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/src/utils"
)

type Event struct {
	ID     uuid.UUID
	From   uuid.UUID
	To     uuid.UUID
	Type   utils.EventType
	Action string
	Data   map[interface{}]interface{}
}

// NewEvent Creates a new event
func NewEvent(from uuid.UUID, to uuid.UUID, eventType utils.EventType, action string, data map[interface{}]interface{}) *Event {
	var null uuid.UUID
	var d = data

	if from == null || to == null {
		return nil
	}

	if d == nil {
		d = make(map[interface{}]interface{}, 0)
	}

	return &Event{
		ID:     uuid.New(),
		From:   from,
		To:     to,
		Type:   eventType,
		Data:   d,
		Action: action,
	}
}
