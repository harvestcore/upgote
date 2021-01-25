package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harvestcore/HarvestCCode/db"
)

// Data Data endpoints.
func Data(router *mux.Router) {
	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Database string `json:"database"`
			Quantity int    `json:"quantity"`
		}

		var request Request
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		if request.Database != "" {
			item := &db.Item{CollectionName: request.Database}
			response := item.Find(make(map[string]interface{}))
			payload, _ = json.Marshal(response)

			if request.Quantity != 0 {
				if len(response.Items) > 0 && request.Quantity > 0 && request.Quantity <= response.Length {
					response.Items = response.Items[response.Length-request.Quantity:]
					response.Length = request.Quantity
					payload, _ = json.Marshal(response)
				} else {
					payload, _ = json.Marshal(map[string]interface{}{"status": false, "message": "Missing database."})
					w.WriteHeader(http.StatusUnprocessableEntity)
				}
			}
		} else {
			payload, _ = json.Marshal(map[string]interface{}{"status": false, "message": "Missing database."})
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("POST")

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Database string `json:"database"`
			Force    bool   `json:"force"`
		}

		var request Request
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		if request.Database != "" && request.Force == true {
			item := &db.Item{CollectionName: request.Database}
			item.Drop()
			payload, _ = json.Marshal(map[string]interface{}{"status": true, "message": "Database removed."})
		} else {
			payload, _ = json.Marshal(map[string]interface{}{"status": false, "message": "Missing database or removal not forced."})
			w.WriteHeader(http.StatusUnprocessableEntity)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("DELETE")
}
