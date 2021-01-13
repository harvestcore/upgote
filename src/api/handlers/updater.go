package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/harvestcore/HarvestCCode/src/core"
	"github.com/harvestcore/HarvestCCode/src/updater"
)

// Updater Updater endpoints.
func Updater(router *mux.Router) {
	router.HandleFunc("/updater", func(w http.ResponseWriter, r *http.Request) {
		c := core.GetCore()
		data := make([]map[string]interface{}, 0)

		for _, value := range c.Updaters {
			data = append(data, map[string]interface{}{
				"database": value.Collection,
				"updater":  &value.Reference,
			})
		}

		payload, _ := json.Marshal(map[string]interface{}{
			"items":  data,
			"status": true,
			"length": len(data),
		})

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("GET")

	router.HandleFunc("/updater", func(w http.ResponseWriter, r *http.Request) {
		var request updater.Updater
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		// TODO ADD MORE CHECKS
		if request.Source != "" {
			c := core.GetCore()
			c.CreateUpdater(map[string]interface{}{
				"schema":      request.Schema,
				"interval":    request.Interval,
				"source":      request.Source,
				"method":      request.Method,
				"requestBody": request.RequestBody,
				"timeout":     request.Timeout,
			})

			payload, _ = json.Marshal(map[string]interface{}{
				"status":  true,
				"message": "Updater created.",
			})
		} else {
			payload, _ = json.Marshal(map[string]interface{}{
				"status":  false,
				"message": "Missing updater ID.",
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("POST")

	router.HandleFunc("/updater", func(w http.ResponseWriter, r *http.Request) {
		var request updater.Updater
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		if request.ID != uuid.Nil {
			c := core.GetCore()
			c.UpdateUpdater(request.ID, map[string]interface{}{
				"schema":      request.Schema,
				"interval":    request.Interval,
				"source":      request.Source,
				"method":      request.Method,
				"requestBody": request.RequestBody,
				"timeout":     request.Timeout,
			})

			payload, _ = json.Marshal(map[string]interface{}{
				"status":  true,
				"message": "Updater updated.",
			})
		} else {
			payload, _ = json.Marshal(map[string]interface{}{
				"status":  false,
				"message": "Missing updater ID.",
			})

			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("PUT")

	router.HandleFunc("/updater", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			ID    string `json:"id"`
			Force bool   `json:"force"`
		}

		var request Request
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		if request.ID != "" && request.Force == true {
			c := core.GetCore()
			id, err := uuid.Parse(request.ID)
			if err != nil {
				payload, _ = json.Marshal(map[string]interface{}{"status": false, "message": "Invalid ID."})
			} else {
				c.StopUpdater(id)
				payload, _ = json.Marshal(map[string]interface{}{"status": true, "message": "Updater removed."})
			}
		} else {
			payload, _ = json.Marshal(map[string]interface{}{"status": false, "message": "Missing updater ID or removal not forced."})
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("DELETE")
}
