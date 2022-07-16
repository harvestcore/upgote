package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/harvestcore/upgote/config"
	"github.com/harvestcore/upgote/core"
	"github.com/harvestcore/upgote/types"
)

// Healthcheck Healthcheck endpoint.
func Healthcheck(router *mux.Router) {
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		data := make(types.Dict)
		data["status"] = core.GetCore() != nil
		data["version"] = config.Get(config.UPGOTE_VERSION).(string)
		payload, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("GET")
}

// Status Status endpoint.
func Status(router *mux.Router) {
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		data := make(types.Dict)

		c := core.GetCore()

		if c == nil {
			data["status"] = false
			data["updaters"] = 0
			data["version"] = config.Get(config.UPGOTE_VERSION).(string)
		} else {
			data["status"] = true
			data["updaters"] = len(c.Updaters)
			data["version"] = config.Get(config.UPGOTE_VERSION).(string)
		}

		payload, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("GET")
}
