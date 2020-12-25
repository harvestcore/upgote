package updater

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/harvestcore/HarvestCCode/src/log"
)

type Updater struct {
	Schema      map[string]interface{}
	Interval    int
	Source      string
	ID          uuid.UUID
	Method      string
	RequestBody map[string]interface{}
	Timeout     int
	Scheduler   *gocron.Scheduler
}

// NewUpdater Creates a new Updater
func NewUpdater(schema map[string]interface{}, interval int, source string, method string, requestBody map[string]interface{}, timeout int) *Updater {
	if len(schema) == 0 {
		log.AddSimple(log.Error, "Updater schema is empty.")
		return nil
	}

	var _interval int = 5
	if interval > _interval {
		_interval = interval
	}

	url, err := url.ParseRequestURI(source)
	if err != nil {
		log.AddSimple(log.Error, "Updater URL is incorrect.")

		return nil
	}

	if method != "GET" && method != "POST" {
		log.AddSimple(log.Error, "Updater method is not GET or POST.")

		return nil
	}

	var _timeout int = 15
	if timeout > _timeout {
		_timeout = timeout
	}

	return &Updater{
		Schema:      schema,
		Interval:    _interval,
		Source:      url.String(),
		ID:          uuid.New(),
		Method:      method,
		RequestBody: requestBody,
		Timeout:     _timeout,
	}
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

					if unmarshalErr != nil {
						log.Add(log.Error, "Data parsing error. ["+u.Method+" "+u.Source+"]", u.ID, uuid.Nil)
					}

					return dat
				}
			} else {
				log.Add(log.Error, "Request response error. ["+u.Method+" "+u.Source+"]", u.ID, uuid.Nil)
			}
		} else {
			log.Add(log.Error, "Request creation error. ["+u.Method+" "+u.Source+"]", u.ID, uuid.Nil)
		}
	} else {
		log.Add(log.Error, "Body parsing error. [POST "+u.Source+"]", u.ID, uuid.Nil)
	}

	return dat
}

// Run Create the scheduler and start running the background taks
func (u *Updater) Run() {
	log.Add(log.Info, "Running updater.", u.ID, uuid.Nil)

	u.Scheduler = gocron.NewScheduler(time.UTC)
	u.Scheduler.StartAsync()

	u.Scheduler.Every(uint64(u.Interval)).Seconds().Do(u.FetchData)
}

// Stop Clears all the background tasks
func (u *Updater) Stop() {
	u.Scheduler.Clear()
}
