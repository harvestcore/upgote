package updater

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Updater struct {
	Schema      map[interface{}]interface{}
	Interval    int
	Source      string
	ID          uuid.UUID
	Method      string
	RequestBody map[string]interface{}
	Timeout     int
}

func (u Updater) SendUpdate() {

}

func (u Updater) HandleEvent() {

}

func (u Updater) GetClient() (client *http.Client) {
	timeout := time.Duration(u.Timeout) * time.Second

	client = &http.Client{
		Timeout: timeout,
	}

	return
}

// FetchData Fetches the data from the source
func (u Updater) FetchData() map[string]interface{} {
	var requestBody []byte
	var requestBodyErr error
	var body bytes.Buffer
	var dat map[string]interface{}

	if u.Method == "POST" {
		requestBody, requestBodyErr = json.Marshal(u.RequestBody)
		body = *bytes.NewBuffer(requestBody)
	}

	if requestBodyErr == nil {
		request, requestErr := http.NewRequest(u.Method, u.Source, &body)

		if requestErr == nil {
			request.Header.Set("Content-type", "application/json")

			response, responseErr := u.GetClient().Do(request)

			if responseErr == nil {
				defer response.Body.Close()

				responseBody, responseBodyErr := ioutil.ReadAll(response.Body)

				if responseBodyErr == nil {
					unmarshalErr := json.Unmarshal(responseBody, &dat)

					if unmarshalErr == nil {
						return dat
					}
				}
			}
		}
	}

	return dat
}

func (u Updater) Run() {

}
