package updater_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/upgote/types"
	"github.com/harvestcore/upgote/updater"
)

func TestUpdaterCreation(t *testing.T) {
	var nilSchema types.Dict
	var nilInterval int
	var nilSource string
	var nilMethod string
	var nilRequestBody types.Dict
	var nilTimeout int

	var schema = make(types.Dict)
	schema["cool"] = "test"

	var u = updater.NewUpdater(
		schema,
		10,
		"https://google.es",
		"GET",
		schema,
		20,
		"myDatabase",
	)
	assert.NotNil(t, u, "Updater creation failed with all parameters set")
	assert.Equal(t, u.Schema, schema, "Mismatch schema")
	assert.Equal(t, u.Interval, 10, "Mismatch interval")
	assert.Equal(t, u.Source, "https://google.es", "Mismatch source")
	assert.Equal(t, u.Method, "GET", "Mismatch method")
	assert.Equal(t, u.RequestBody, schema, "Mismatch requestBody")
	assert.Equal(t, u.Timeout, 20, "Mismatch timeout")
	assert.Equal(t, u.Collection, "myDatabase", "Mismatch collection")

	u = updater.NewUpdater(
		nilSchema,
		10,
		"https://google.es",
		"GET",
		schema,
		20,
		"myDatabase",
	)
	assert.Nil(t, u, "Updater created without schema")

	u = updater.NewUpdater(
		schema,
		nilInterval,
		"https://google.es",
		"GET",
		schema,
		20,
		"myDatabase",
	)
	assert.NotNil(t, u, "Updater not created without interval")
	assert.Equal(t, u.Interval, 5, "Interval is not 5")

	u = updater.NewUpdater(
		schema,
		10,
		nilSource,
		"GET",
		schema,
		20,
		"myDatabase",
	)
	assert.Nil(t, u, "Updater created without interval")

	u = updater.NewUpdater(
		schema,
		10,
		"https://google.es",
		nilMethod,
		schema,
		20,
		"myDatabase",
	)
	assert.Nil(t, u, "Updater created without method")

	u = updater.NewUpdater(
		schema,
		10,
		"https://google.es",
		"GET",
		nilRequestBody,
		20,
		"myDatabase",
	)
	assert.NotNil(t, u, "Updater not created without requestBody")

	u = updater.NewUpdater(
		schema,
		10,
		"https://google.es",
		"GET",
		schema,
		nilTimeout,
		"myDatabase",
	)
	assert.NotNil(t, u, "Updater not created without timeout")
	assert.Equal(t, u.Timeout, 15, "Timeout is not 15")
}

func TestUpdaterUpdate(t *testing.T) {
	var schema = make(types.Dict)
	schema["cool"] = "test"

	var u = updater.NewUpdater(
		schema,
		10,
		"https://google.es",
		"GET",
		schema,
		20,
		"myDatabase",
	)

	schema["cool"] = "test2"
	schema["test2"] = "cool"

	var data = make(types.Dict)
	data["schema"] = schema
	data["interval"] = 60
	data["source"] = "https://google.com"
	data["method"] = "POST"
	data["requestBody"] = make(types.Dict)
	data["timeout"] = 30

	u.Update(data)
}
