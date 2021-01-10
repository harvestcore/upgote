package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harvestcore/HarvestCCode/src/core"
	"github.com/harvestcore/HarvestCCode/src/handler"
)

// Healthcheck Healthcheck endpoint.
func Healthcheck(router *mux.Router) {
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["status"] = core.GetCore() != nil && handler.GetHandler() != nil
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
		h := handler.GetHandler()

		if c == nil || h == nil {
			data["status"] = false
		} else {
			data["status"] = true
			data["updaters"] = len(c.Updaters)
			data["events"] = len(h.EventQueue)
		}

		payload, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("GET")
}
