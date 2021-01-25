package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/harvestcore/HarvestCCode/api/tests"
	"github.com/harvestcore/HarvestCCode/db"
	"github.com/harvestcore/HarvestCCode/log"
	"github.com/harvestcore/HarvestCCode/utils"
)

func TestPostDataNoArgs(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostDataNoArgs")
		req, _ := http.NewRequest("POST", "/api/data", bytes.NewBuffer([]byte(`{}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "POST /data (no args) status code is not 422")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.False(t, data["status"].(bool), "POST /data (no args) status is not false")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostDataNoArgs")
	}
}

func TestPostDataWithDatabaseNoQuantity(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostDataWithDatabaseNoQuantity")
		req, _ := http.NewRequest("POST", "/api/data", bytes.NewBuffer([]byte(`{"database": "log"}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "POST /data status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "POST /data status is not true")
		assert.GreaterOrEqual(t, data["length"], 0.0, "POST /data length is not valid")
		assert.GreaterOrEqual(t, len(data["items"].([]interface{})), 0, "POST /data items are not valid")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostDataWithDatabaseNoQuantity")
	}
}

func TestPostDataWithWrongDatabaseNoQuantity(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostDataWithWrongDatabaseNoQuantity")
		req, _ := http.NewRequest("POST", "/api/data", bytes.NewBuffer([]byte(`{"database": "__test__"}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "POST /data status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "POST /data status is not true")
		assert.Equal(t, data["length"], 0.0, "POST /data length is not valid")
		assert.Equal(t, len(data["items"].([]interface{})), 0, "POST /data items are not valid")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostDataWithWrongDatabaseNoQuantity")
	}
}

func TestPostDataWithDatabaseAndQuantity(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostDataWithDatabaseAndQuantity")
		item := &db.Item{CollectionName: "TestPostDataWithDatabaseAndQuantity"}
		item.InsertOne(map[string]interface{}{"test": 1})

		req, _ := http.NewRequest("POST", "/api/data", bytes.NewBuffer([]byte(`{"database": "TestPostDataWithDatabaseAndQuantity", "quantity": 1}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusOK, "POST /data status code is not 200")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "POST /data status is not true")
		assert.Equal(t, data["length"], 1.0, "POST /data length is not valid")
		assert.Equal(t, len(data["items"].([]interface{})), 1, "POST /data items are not valid")

		item.Drop()
		log.AddSimple(log.Info, "@TEST-END # Running TestPostDataWithDatabaseAndQuantity")
	}
}

func TestPostDataWithDatabaseAndWrongQuantity(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestPostDataWithDatabaseAndWrongQuantity")
		req, _ := http.NewRequest("POST", "/api/data", bytes.NewBuffer([]byte(`{"database": "log", "quantity": -1}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "POST /data status code is not 422")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.False(t, data["status"].(bool), "POST /data status is not false")
		assert.NotEqual(t, data["message"], "", "POST /data message is empty")
		log.AddSimple(log.Info, "@TEST-END # Running TestPostDataWithDatabaseAndWrongQuantity")
	}
}

func TestDeleteDataWithNoDatabase(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestDeleteDataWithNoDatabase")
		req, _ := http.NewRequest("DELETE", "/api/data", bytes.NewBuffer([]byte(`{}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "DELETE /data status code is not 422")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "DELETE /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.False(t, data["status"].(bool), "DELETE /data status is not false")
		assert.NotEqual(t, data["message"], "", "DELETE /data message is empty")
		log.AddSimple(log.Info, "@TEST-END # Running TestDeleteDataWithNoDatabase")
	}
}

func TestDeleteDataWithWrongDatabaseAndNoForce(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestDeleteDataWithWrongDatabaseAndNoForce")
		req, _ := http.NewRequest("DELETE", "/api/data", bytes.NewBuffer([]byte(`{"database": "yikes"}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "DELETE /data status code is not 422")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "DELETE /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.False(t, data["status"].(bool), "DELETE /data status is not false")
		assert.NotEqual(t, data["message"], "", "DELETE /data message is empty")
		log.AddSimple(log.Info, "@TEST-END # Running TestDeleteDataWithWrongDatabaseAndNoForce")
	}
}

func TestDeleteDataWithWrongDatabaseAndForce(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestDeleteDataWithWrongDatabaseAndForce")
		req, _ := http.NewRequest("DELETE", "/api/data", bytes.NewBuffer([]byte(`{"database": "yikes", "force": true}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusNoContent, "DELETE /data status code is not 204")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "DELETE /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "DELETE /data status is not false")
		assert.NotEqual(t, data["message"], "", "DELETE /data message is empty")
		log.AddSimple(log.Info, "@TEST-END # Running TestDeleteDataWithWrongDatabaseAndForce")
	}
}

func TestDeleteDataWithExistingDatabaseAndForce(t *testing.T) {
	if !utils.RunningInDocker() {
		log.AddSimple(log.Info, "@TEST # Running TestDeleteDataWithExistingDatabaseAndForce")
		item := &db.Item{CollectionName: "__delete"}
		item.InsertOne(map[string]interface{}{"test": 1})

		req, _ := http.NewRequest("DELETE", "/api/data", bytes.NewBuffer([]byte(`{"database": "__delete", "force": true}`)))
		res := api.ExecuteTestingRequest(req)

		assert.Equal(t, res.Code, http.StatusNoContent, "DELETE /data status code is not 204")
		assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "DELETE /data Content type is not \"application/json\"")

		var data map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &data)

		assert.True(t, data["status"].(bool), "DELETE /data status is not false")
		assert.NotEqual(t, data["message"], "", "DELETE /data message is empty")
		log.AddSimple(log.Info, "@TEST-END # Running TestDeleteDataWithExistingDatabaseAndForce")
	}
}
