package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	api "github.com/harvestcore/HarvestCCode/src/api/tests"
)

func TestGetLogFile(t *testing.T) {
	req, _ := http.NewRequest("GET", "/log", nil)
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "GET /log status code is not 200")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "text/plain", "GET /log Content type is not \"text/plain\"")
}

func TestPostLogFileNoQuantity(t *testing.T) {
	req, _ := http.NewRequest("POST", "/log", bytes.NewBuffer([]byte(`{}`)))
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "GET /log status code is not 200. Without quantity")
	assert.Equal(t, res.HeaderMap.Get("Content-Type"), "application/json", "POST /log Content type is not \"application/json\"")

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "POST /log status is not true")
	assert.Positive(t, data["length"], "POST /log length is not valid")
	assert.Positive(t, len(data["items"].([]interface{})), "POST /log items are not valid")

}

func TestPostLogFileWithQuantity(t *testing.T) {
	req, _ := http.NewRequest("POST", "/log", bytes.NewBuffer([]byte(`{"quantity": 1}`)))
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusOK, "GET /log status code is not 200. With quantity")

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.True(t, data["status"].(bool), "POST /log status is not true")
	assert.Equal(t, data["length"], 1.0, "POST /log length is not 1")
	assert.Equal(t, len(data["items"].([]interface{})), 1, "POST /log items are not 1")
}

func TestPostLogFileNegativeQuantity(t *testing.T) {
	req, _ := http.NewRequest("POST", "/log", bytes.NewBuffer([]byte(`{"quantity": -1}`)))
	res := api.ExecuteTestingRequest(req)

	assert.Equal(t, res.Code, http.StatusUnprocessableEntity, "GET /log status code is not 422. With negative quantity")

	var data map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.False(t, data["status"].(bool), "POST /log status is not false. With negative quantity")
	assert.NotEqual(t, data["message"], "", "POST /log message is empty")
}
