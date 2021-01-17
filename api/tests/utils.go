package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/harvestcore/HarvestCCode/api"
)

// ExecuteTestingRequest Executes the given request in the testing router
func ExecuteTestingRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	api.GetServer().Server.Handler.ServeHTTP(recorder, req)

	return recorder
}
