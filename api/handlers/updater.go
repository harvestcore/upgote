package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/harvestcore/HarvestCCode/core"
)

type customRequest struct {
	Schema      map[string]interface{} `json:"schema"`
	Interval    int                    `json:"interval"`
	Source      string                 `json:"source"`
	ID          uuid.UUID              `json:"id"`
	Method      string                 `json:"method"`
	RequestBody map[string]interface{} `json:"requestBody"`
	Timeout     int                    `json:"timeout"`
	Collection  string                 `json:"database"`
}

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
		var request customRequest
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		if checkUpdaterParams(request) {
			c := core.GetCore()
			updaterID := c.CreateUpdater(map[string]interface{}{
				"schema":      request.Schema,
				"interval":    request.Interval,
				"source":      request.Source,
				"method":      request.Method,
				"requestBody": request.RequestBody,
				"timeout":     request.Timeout,
				"collection":  request.Collection,
			})

			if updaterID != uuid.Nil {
				payload, _ = json.Marshal(map[string]interface{}{
					"status":  true,
					"message": "Updater created.",
					"id":      updaterID.String(),
				})
			} else {
				payload, _ = json.Marshal(map[string]interface{}{
					"status":  false,
					"message": "Error creating updater. Wrong parameters",
				})
				w.WriteHeader(http.StatusUnprocessableEntity)
			}
		} else {
			payload, _ = json.Marshal(map[string]interface{}{
				"status":  false,
				"message": "Wrong parameters",
			})
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("POST")

	router.HandleFunc("/updater", func(w http.ResponseWriter, r *http.Request) {
		var request customRequest
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		if request.ID != uuid.Nil || checkUpdaterParams(request) {
			c := core.GetCore()
			c.UpdateUpdater(request.ID, map[string]interface{}{
				"schema":      request.Schema,
				"interval":    request.Interval,
				"source":      request.Source,
				"method":      request.Method,
				"requestBody": request.RequestBody,
				"timeout":     request.Timeout,
				"collection":  request.Collection,
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
				w.WriteHeader(http.StatusUnprocessableEntity)
			} else {
				c.RemoveUpdater(id)
				payload, _ = json.Marshal(map[string]interface{}{"status": true, "message": "Updater removed."})
				w.WriteHeader(http.StatusNoContent)
			}
		} else {
			payload, _ = json.Marshal(map[string]interface{}{"status": false, "message": "Missing updater ID or removal not forced."})
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("DELETE")

	router.HandleFunc("/updater/action", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			ID     string `json:"id"`
			Action string `json:"action"`
		}

		var payload []byte
		var request Request

		json.NewDecoder(r.Body).Decode(&request)

		if request.ID != "" || (request.Action != "start" && request.Action != "stop") {
			parsedID, _ := uuid.Parse(request.ID)

			if parsedID != uuid.Nil {
				c := core.GetCore()
				status := false

				if request.Action == "start" {
					status = c.StartUpdater(parsedID)
				} else if request.Action == "stop" {
					status = c.StopUpdater(parsedID)
				}

				if status {
					payload, _ = json.Marshal(map[string]interface{}{
						"status":  true,
						"message": "Action executed successfully.",
					})
				} else {
					payload, _ = json.Marshal(map[string]interface{}{
						"status":  false,
						"message": "The updater does not exist.",
					})
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			} else {
				payload, _ = json.Marshal(map[string]interface{}{
					"status":  false,
					"message": "Wrong ID.",
				})
				w.WriteHeader(http.StatusUnprocessableEntity)
			}
		} else {
			payload, _ = json.Marshal(map[string]interface{}{
				"status":  false,
				"message": "Wrong fields.",
			})
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("POST")
}

func checkUpdaterParams(updater customRequest) bool {
	if len(updater.Schema) == 0 ||
		updater.Collection == "" ||
		updater.Schema == nil ||
		updater.Interval <= 0 ||
		updater.Source == "" ||
		updater.Method == "" ||
		updater.Timeout <= 0 {
		return false
	}

	return true
}
