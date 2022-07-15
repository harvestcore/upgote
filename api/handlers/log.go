package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harvestcore/upgote/config"
	"github.com/harvestcore/upgote/log"
)

// Log Log endpoint.
func Log(router *mux.Router) {
	router.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		logFilePath := config.GetManager().GetVariable(config.LOG_FILE)
		file, _ := ioutil.ReadFile(logFilePath)

		w.Header().Set("Content-Type", "text/plain")
		w.Write(file)
	}).Methods("GET")

	router.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Quantity int `json:"quantity"`
		}

		var request Request
		var payload []byte

		json.NewDecoder(r.Body).Decode(&request)

		response := log.GetLogger().Item.Find(make(map[string]interface{}))

		payload, _ = json.Marshal(response)

		if request.Quantity != 0 {
			if request.Quantity > 0 && request.Quantity <= response.Length {
				response.Items = response.Items[response.Length-request.Quantity:]
				response.Length = request.Quantity

				payload, _ = json.Marshal(response)
			} else {
				payload, _ = json.Marshal(map[string]interface{}{
					"status":  false,
					"message": "Wrong quantity value.",
				})

				w.WriteHeader(http.StatusUnprocessableEntity)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}).Methods("POST")
}
