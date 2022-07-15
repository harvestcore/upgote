package core_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	
	"github.com/harvestcore/upgote/core"
	"github.com/harvestcore/upgote/utils"
)

func TestCoreCreation(t *testing.T) {
	var c = core.GetCore()

	assert.NotNil(t, c, "Core creation failed")
	assert.NotEqual(t, c.ID, uuid.Nil, "Mismatch ID")
	assert.Equal(t, len(c.Updaters), 0, "Updaters map is not empty")
}

func TestCoreCreateUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		var schema = make(map[string]interface{})
		schema["cool"] = "test"

		var data = make(map[string]interface{})
		data["collection"] = "testing"
		data["schema"] = schema
		data["interval"] = 10
		data["source"] = "https://google.es"
		data["method"] = "GET"
		data["requestBody"] = make(map[string]interface{})
		data["timeout"] = 20

		var c = core.GetCore()

		var updater = c.CreateUpdater(data)
		assert.NotEqual(t, updater, uuid.Nil, "Updater creation via Core failed")
		assert.Equal(t, len(c.Updaters), 1, "There is more than one updater created")

		c.Updaters = make(map[uuid.UUID]*core.UpdaterMap)
	}
}

func TestCoreStopUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		var schema = make(map[string]interface{})
		schema["cool"] = "test"

		var data = make(map[string]interface{})
		data["collection"] = "yikes"
		data["schema"] = schema
		data["interval"] = 10
		data["source"] = "https://google.es"
		data["method"] = "GET"
		data["requestBody"] = make(map[string]interface{})
		data["timeout"] = 20

		var c = core.GetCore()

		var updater = c.CreateUpdater(data)
		assert.NotEqual(t, updater, uuid.Nil, "Updater creation via Core failed")
		assert.Equal(t, len(c.Updaters), 1, "There is more than one updater created")

		c.StopUpdater(updater)
		assert.Equal(t, len(c.Updaters), 1, "Updater not stopped")

		c.RemoveUpdater(updater)
	}
}

func TestCoreRemoveUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		var schema = make(map[string]interface{})
		schema["cool"] = "test"

		var data = make(map[string]interface{})
		data["collection"] = "yikes123"
		data["schema"] = schema
		data["interval"] = 10
		data["source"] = "https://google.es"
		data["method"] = "GET"
		data["requestBody"] = make(map[string]interface{})
		data["timeout"] = 20

		var c = core.GetCore()

		var updater = c.CreateUpdater(data)
		assert.NotEqual(t, updater, uuid.Nil, "Updater creation via Core failed")
		assert.Equal(t, len(c.Updaters), 1, "There is more than one updater created")

		c.RemoveUpdater(updater)
		assert.Equal(t, len(c.Updaters), 0, "Updater not stopped")
	}
}
