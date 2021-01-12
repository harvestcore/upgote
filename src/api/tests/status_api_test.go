package api_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/harvestcore/HarvestCCode/src/api/tests"
)

func TestGetHeartcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "GET /healthcheck status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "GET /healthcheck Content type is not \"application/json\"")

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "GET /healthcheck status is not true")
}

func TestGetStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "/status", nil)
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "GET /status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "GET /status Content type is not \"application/json\"")

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "GET /status status is not true")
	assert.GreaterOrEqual(t, data["updaters"].(float64), 0.0, "GET /status wrong updaters")
	assert.GreaterOrEqual(t, data["events"].(float64), 0.0, "GET /status wrong events")
}
