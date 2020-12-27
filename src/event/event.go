package event

import (
	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/src/log"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

type Event struct {
	ID     uuid.UUID
	From   uuid.UUID
	To     uuid.UUID
	Type   utils.EventType
	Action string
	Data   map[string]interface{}
}

// NewEvent Creates a new event
func NewEvent(from uuid.UUID, to uuid.UUID, eventType utils.EventType, action string, data map[string]interface{}) *Event {
	var null uuid.UUID
	var d = data

	if from == null || to == null {
		return nil
	}

	var id = uuid.New()

	if d == nil {
		d = make(map[string]interface{}, 0)

		log.AddSimple(log.Warning, "Created event "+id.String()+" without data.")
	}

	return &Event{
		ID:     id,
		From:   from,
		To:     to,
		Type:   eventType,
		Data:   d,
		Action: action,
	}
}
