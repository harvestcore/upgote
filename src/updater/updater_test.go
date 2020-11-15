package updater_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/HarvestCCode/src/updater"
)

func TestUpdaterCreation(t *testing.T) {
	var nilSchema map[string]interface{}
	var nilInterval int
	var nilSource string
	var nilMethod string
	var nilRequestBody map[string]interface{}
	var nilTimeout int

	var schema = make(map[string]interface{})
	schema["cool"] = "test"

	var u = updater.New(
		schema,
		10,
		"https://google.es",
		"GET",
		schema,
		20,
	)
	assert.NotNil(t, u, "Updater creation failed with all parameters set")
	assert.Equal(t, u.Schema, schema, "Mismatch schema")
	assert.Equal(t, u.Interval, 10, "Mismatch interval")
	assert.Equal(t, u.Source, "https://google.es", "Mismatch source")
	assert.Equal(t, u.Method, "GET", "Mismatch method")
	assert.Equal(t, u.RequestBody, schema, "Mismatch requestBody")
	assert.Equal(t, u.Timeout, 20, "Mismatch timeout")

	u = updater.New(
		nilSchema,
		10,
		"https://google.es",
		"GET",
		schema,
		20,
	)
	assert.Nil(t, u, "Updater created without schema")

	u = updater.New(
		schema,
		nilInterval,
		"https://google.es",
		"GET",
		schema,
		20,
	)
	assert.NotNil(t, u, "Updater not created without interval")
	assert.Equal(t, u.Interval, 5, "Interval is not 5")

	u = updater.New(
		schema,
		10,
		nilSource,
		"GET",
		schema,
		20,
	)
	assert.Nil(t, u, "Updater created without interval")

	u = updater.New(
		schema,
		10,
		"https://google.es",
		nilMethod,
		schema,
		20,
	)
	assert.Nil(t, u, "Updater created without method")

	u = updater.New(
		schema,
		10,
		"https://google.es",
		"GET",
		nilRequestBody,
		20,
	)
	assert.NotNil(t, u, "Updater not created without requestBody")

	u = updater.New(
		schema,
		10,
		"https://google.es",
		"GET",
		schema,
		nilTimeout,
	)
	assert.NotNil(t, u, "Updater not created without timeout")
	assert.Equal(t, u.Timeout, 15, "Timeout is not 15")
}
