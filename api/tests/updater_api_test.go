package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/harvestcore/HarvestCCode/api/tests"
	"github.com/harvestcore/HarvestCCode/log"
	"github.com/harvestcore/HarvestCCode/utils"
)

func TestGetUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestGetUpdater")
		req, _ := http.NewRequest("GET", "/updater", nil)
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "GET /updater status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "GET /updater Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "GET /updater status is not true")
		assert.GreaterOrEqual(t, len(data["items"].([]interface{})), 0, "GET /updater items are invalid")
		assert.GreaterOrEqual(t, data["length"], 0.0, "GET /updater length is invalid")
		log.AddSimple(log.Info, "@TEST-END # Running TestGetUpdater")
	}
}

func TestPostCreateUpdaterWrongParams(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostCreateUpdaterWrongParams")
		req, _ := http.NewRequest("POST", "/updater", bytes.NewBuffer([]byte(`{"interval": -4, "method": "DELETE"}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "POST /updater status code is not 422")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /updater Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.False(t, data["status"].(bool), "POST /updater status is not false")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostCreateUpdaterWrongParams")
	}
}

func TestDeleteUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestDeleteUpdater")
		req, _ := http.NewRequest("POST", "/updater", bytes.NewBuffer([]byte(`{"database": "testingDELETE", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
		res := api.ExecuteTestingRequest(req)

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		req, _ = http.NewRequest("DELETE", "/updater", bytes.NewBuffer([]byte(`{"force": true, "id": "`+data["id"].(string)+`"}`)))
		res = api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusNoContent, "DELETE /updater status code is not 204")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /updater Content type is not \"application/json\"")

		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "DELETE /updater status is not true")
		log.AddSimple(log.Info, "@TEST-END # Running TestDeleteUpdater")
	}
}

func TestPostCreateUpdaterParams(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostCreateUpdaterParams")
		req, _ := http.NewRequest("POST", "/updater", bytes.NewBuffer([]byte(`{"database": "testingPOST", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "POST /updater status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /updater Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "POST /updater status is not true")

		req, _ = http.NewRequest("DELETE", "/updater", bytes.NewBuffer([]byte(`{"force": true, "id": "`+data["id"].(string)+`"}`)))
		res = api.ExecuteTestingRequest(req)
		log.AddSimple(log.Info, "@TEST-END # Running TestPostCreateUpdaterParams")
	}
}

func TestPutUpdateUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPutUpdateUpdater")
		req, _ := http.NewRequest("POST", "/updater", bytes.NewBuffer([]byte(`{"database": "testingPUT", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
		res := api.ExecuteTestingRequest(req)

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		req, _ = http.NewRequest("PUT", "/updater", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
		res = api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "PUT /updater status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "PUT /updater Content type is not \"application/json\"")

		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "PUT /updater status is not true")
		log.AddSimple(log.Info, "@TEST-END # Running TestPutUpdateUpdater")
	}
}

func TestPostStartUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostStartUpdater")
		req, _ := http.NewRequest("POST", "/updater", bytes.NewBuffer([]byte(`{"database": "testingPOST", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
		res := api.ExecuteTestingRequest(req)

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		req, _ = http.NewRequest("POST", "/updater/start", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`"}`)))
		res = api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "POST /updater/start status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "PUT /updater Content type is not \"application/json\"")

		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "POST /updater/start status is not true")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostStartUpdater")
	}
}

func TestPostStopUpdater(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostStopUpdater")
		req, _ := http.NewRequest("POST", "/updater", bytes.NewBuffer([]byte(`{"database": "testingPOSTstop", "schema": {"my": "schema"}, "interval": 60, "source": "https://ipinfo.io/json", "method": "GET", "timeout": 30}`)))
		res := api.ExecuteTestingRequest(req)

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		req, _ = http.NewRequest("POST", "/updater/start", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`"}`)))
		res = api.ExecuteTestingRequest(req)

		req, _ = http.NewRequest("POST", "/updater/stop", bytes.NewBuffer([]byte(`{"id": "`+data["id"].(string)+`"}`)))
		res = api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "POST /updater/start status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "PUT /updater Content type is not \"application/json\"")

		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "POST /updater/start status is not true")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostStopUpdater")
	}
}
