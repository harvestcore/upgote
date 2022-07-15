package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harvestcore/upgote/core"
)

// Healthcheck Healthcheck endpoint.
func Healthcheck(router *mux.Router) {
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["status"] = core.GetCore() != nil
		payload, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("GET")
}

// Status Status endpoint.
func Status(router *mux.Router) {
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})

		c := core.GetCore()

		if c == nil {
			data["status"] = false
		} else {
			data["status"] = true
			data["updaters"] = len(c.Updaters)
		}

		payload, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("GET")
}
