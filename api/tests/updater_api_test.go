package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/harvestcore/upgote/api/tests"
	"github.com/harvestcore/upgote/types"
)

func TestGetUpdater(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/updater", nil)
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "GET /updater status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "GET /updater Content type is not \"application/json\"")

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "GET /updater status is not true")
	assert.GreaterOrEqual(t, len(data["items"].([]interface{})), 0, "GET /updater items are invalid")
	assert.GreaterOrEqual(t, data["length"], 0.0, "GET /updater length is invalid")
}

func TestPostCreateUpdaterWrongParams(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/updater", bytes.NewBuffer([]byte(`{"interval": -4, "method": "DELETE"}`)))
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "POST /updater status code is not 422")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /updater Content type is not \"application/json\"")

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.False(t, data["status"].(bool), "POST /updater status is not false")
}

func TestDeleteUpdater(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/updater", bytes.NewBuffer([]byte(`{"database": "testingDELETE", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
	res := api.ExecuteTestingRequest(req)

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	req, _ = http.NewRequest("DELETE", "/api/updater", bytes.NewBuffer([]byte(`{"force": true, "id": "`+data["id"].(string)+`"}`)))
	res = api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "DELETE /updater status code is not 204")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /updater Content type is not \"application/json\"")

	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "DELETE /updater status is not true")
}

func TestPostCreateUpdaterParams(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/updater", bytes.NewBuffer([]byte(`{"database": "testingPOST", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusCreated, "POST /updater status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /updater Content type is not \"application/json\"")

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "POST /updater status is not true")

	req, _ = http.NewRequest("DELETE", "/api/updater", bytes.NewBuffer([]byte(`{"force": true, "id": "`+data["id"].(string)+`"}`)))
	res = api.ExecuteTestingRequest(req)
}

func TestPutUpdateUpdater(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/updater", bytes.NewBuffer([]byte(`{"database": "testingPUT", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
	res := api.ExecuteTestingRequest(req)

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	req, _ = http.NewRequest("PUT", "/api/updater", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
	res = api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "PUT /updater status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "PUT /updater Content type is not \"application/json\"")

	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "PUT /updater status is not true")
}

func TestPostStartUpdater(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/updater", bytes.NewBuffer([]byte(`{"database": "testingPOST", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
	res := api.ExecuteTestingRequest(req)

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	req, _ = http.NewRequest("POST", "/api/updater/action", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`", "action": "start"}`)))
	res = api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "POST /updater/action start status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "PUT /updater Content type is not \"application/json\"")

	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "POST /updater/action start action status is not true")
}

func TestPostStopUpdater(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/updater", bytes.NewBuffer([]byte(`{"database": "testingPOSTstop", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
	res := api.ExecuteTestingRequest(req)

	var data types.Dict
	json.Unmarshal(res.Body.Bytes(), &data)

	req, _ = http.NewRequest("POST", "/api/updater/action", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`", "action": "start"}`)))
	res = api.ExecuteTestingRequest(req)

	req, _ = http.NewRequest("POST", "/api/updater/action", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`", "action": "stop"}`)))
	res = api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "POST /updater/action start status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "PUT /updater Content type is not \"application/json\"")

	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "POST /updater/action stop status is not true")
}
