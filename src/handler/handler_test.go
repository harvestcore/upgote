package handler_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/handler"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

func TestHandlerCreation(t *testing.T) {
	var h = handler.GetHandler()

	assert.NotNil(t, h, "Handler creation failed with all parameters set")
	assert.Equal(t, h.Timeout, 1, "Mismatch timeout")
	assert.Equal(t, h.Lock, false, "Mismatch lock")
	assert.NotNil(t, h.Scheduler, "Scheduler not created")
	assert.NotNil(t, h.EventQueue, "Nil QueueEvent")
	assert.Equal(t, len(h.EventQueue), 0, "QueueEvent is not empty")
}

func TestHandlerSchedulerCreation(t *testing.T) {
	var h = handler.GetHandler()
	assert.NotNil(t, h, "Handler creation failed with all parameters set")

	h.StartHandlingEvents()
	assert.NotNil(t, h.Scheduler, "Nil Scheduler after StartHandlingEvents")
}

func TestHandlerSchedulerEvents(t *testing.T) {
	var h = handler.GetHandler()
	assert.NotNil(t, h, "Handler creation failed")

	h.StartHandlingEvents()
	assert.NotNil(t, h.Scheduler, "Nil Scheduler after StartHandlingEvents")

	e1 := event.NewEvent(
		uuid.New(),
		uuid.New(),
		utils.StoreData,
		"test",
		make(map[string]interface{}, 0),
	)

	e2 := event.NewEvent(
		uuid.New(),
		uuid.New(),
		utils.StoreData,
		"test",
		make(map[string]interface{}, 0),
	)

	h.EventQueue = append(h.EventQueue, *e1)
	h.EventQueue = append(h.EventQueue, *e2)

	assert.Equal(t, len(h.EventQueue), 2, "QueueEvent does not have 2 events")

	h.ClearEventQueue()
	assert.Equal(t, len(h.EventQueue), 0, "QueueEvent is not empty")
}

func TestHandlerSchedulerHandleEvents(t *testing.T) {
	var h = handler.GetHandler()
	assert.NotNil(t, h, "Handler creation failed")

	h.StartHandlingEvents()
	assert.NotNil(t, h.Scheduler, "Nil Scheduler after StartHandlingEvents")

	e1 := event.NewEvent(
		uuid.New(),
		uuid.New(),
		utils.StoreData,
		"test",
		make(map[string]interface{}, 0),
	)

	e2 := event.NewEvent(
		uuid.New(),
		uuid.New(),
		utils.StoreData,
		"test",
		make(map[string]interface{}, 0),
	)

	h.EventQueue = append(h.EventQueue, *e1)
	h.EventQueue = append(h.EventQueue, *e2)
	assert.Equal(t, len(h.EventQueue), 2, "QueueEvent does not have 2 events")

	h.StartHandlingEvents()

	// Wait at least 3 secs (it could be lower but it is set to 3s to make sure)
	// so the Handler can start working
	time.Sleep(time.Duration(3) * time.Second)
	assert.Equal(t, len(h.EventQueue), 0, "QueueEvent is not empty")

	h.StopHandlingEvents()
	assert.Equal(t, len(h.Scheduler.Jobs()), 0, "The scheduler has pending jobs")
}
