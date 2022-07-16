package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/harvestcore/upgote/api"
)

// ExecuteTestingRequest Executes the given request in the testing router.
func ExecuteTestingRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	var server = api.GetServer().Server
	server.Handler.ServeHTTP(recorder, req)

	return recorder
}
