package event_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestEventCreation(t *testing.T) {
	var niluuid uuid.UUID
	var nildata map[string]interface{}

	var from = uuid.New()
	var to = uuid.New()
	var data = make(map[string]interface{}, 0)

	var e = event.NewEvent(
		from,
		to,
		utils.Store,
		"action",
		data,
	)

	assert.NotNil(t, e, "Event creation return nil with all parameters set")
	assert.NotNil(t, e.ID, "ID is nil")
	assert.Equal(t, e.From, from, "Mismatch 'from' UUID")
	assert.Equal(t, e.To, to, "Mismatch 'to' UUID")
	assert.Equal(t, e.Type, utils.EventType("store"), "Mismatch event type")
	assert.Equal(t, e.Action, "action", "Mismatch action")
	assert.NotNil(t, e.Data, "Data map is nil")
	assert.Equal(t, len(e.Data), 0, "Data map length is not 0")

	e = event.NewEvent(
		niluuid,
		to,
		utils.Store,
		"action",
		data,
	)

	assert.Nil(t, e, "Event created with nil 'from'")

	e = event.NewEvent(
		from,
		niluuid,
		utils.Store,
		"action",
		data,
	)

	assert.Nil(t, e, "Event created with nil 'to'")

	e = event.NewEvent(
		from,
		to,
		utils.Store,
		"action",
		nildata,
	)

	assert.NotNil(t, e, "Event not created with nil 'data'")
	assert.NotNil(t, e.Data, "Data map is nil")
	assert.Equal(t, len(e.Data), 0, "Data map length is not 0")
}
