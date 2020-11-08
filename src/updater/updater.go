package updater

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
)

type Updater struct {
	Schema      map[interface{}]interface{}
	Interval    uint64
	Source      string
	ID          uuid.UUID
	Method      string
	RequestBody map[string]interface{}
	Timeout     int
	Scheduler   *gocron.Scheduler
}

func (u *Updater) SendUpdate() {

}

func (u *Updater) HandleEvent() {

}

// GetClient Returns the client configured with the timeout inverval
func (u *Updater) GetClient() (client *http.Client) {
	timeout := time.Duration(u.Timeout) * time.Second

	client = &http.Client{
		Timeout: timeout,
	}

	return
}

// FetchData Fetches the data from the source
func (u *Updater) FetchData() map[string]interface{} {
	var requestBody []byte
	var requestBodyErr error
	var body bytes.Buffer
	var dat map[string]interface{}

	// If the method is POST -> encode the request body
	if u.Method == "POST" {
		requestBody, requestBodyErr = json.Marshal(u.RequestBody)
		body = *bytes.NewBuffer(requestBody)
	}

	if requestBodyErr == nil {
		// Create request object
		request, requestErr := http.NewRequest(u.Method, u.Source, &body)

		if requestErr == nil {
			// Set header content type and perform the request
			request.Header.Set("Content-type", "application/json")
			response, responseErr := u.GetClient().Do(request)

			if responseErr == nil {
				// Defer the close of the body. It will be closed as soon as
				// this method ends
				defer response.Body.Close()

				// Read the body
				responseBody, responseBodyErr := ioutil.ReadAll(response.Body)

				if responseBodyErr == nil {
					// Decode the body
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

// Run Create the scheduler and start running the background taks
func (u *Updater) Run() {
	u.Scheduler = gocron.NewScheduler(time.UTC)
	u.Scheduler.StartAsync()

	u.Scheduler.Every(u.Interval).Seconds().Do(u.FetchData)
}

// Stop Clears all the background tasks
func (u *Updater) Stop() {
	u.Scheduler.Clear()
}
