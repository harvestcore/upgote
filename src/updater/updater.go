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

	"github.com/harvestcore/HarvestCCode/src/db"
	"github.com/harvestcore/HarvestCCode/src/log"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

type Updater struct {
	Schema      map[string]interface{} `json:"schema"`
	Interval    int                    `json:"interval"`
	Source      string                 `json:"source"`
	ID          uuid.UUID              `json:"id"`
	Method      string                 `json:"method"`
	RequestBody map[string]interface{} `json:"requestBody"`
	Timeout     int                    `json:"timeout"`
	Collection  string                 `json:"collection"`

	scheduler *gocron.Scheduler
}

// NewUpdater Creates a new Updater
func NewUpdater(schema map[string]interface{}, interval int, source string, method string, requestBody map[string]interface{}, timeout int, collection string) *Updater {
	if schema == nil {
		log.AddSimple(log.Error, "Updater schema is not valid.")
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

	if collection == "" {
		log.AddSimple(log.Error, "Updater collection is empty")

		return nil
	}

	var _timeout int = 15
	if timeout > _timeout {
		_timeout = timeout
	}

	id := uuid.New()
	var updater *Updater

	updater = &Updater{
		Schema:      schema,
		Interval:    _interval,
		Source:      url.String(),
		ID:          id,
		Method:      method,
		RequestBody: requestBody,
		Timeout:     _timeout,
		Collection:  collection,
	}

	return updater
}

// SendUpdate Stores the fetched data
func (u *Updater) SendUpdate(data map[string]interface{}) {
	item := &db.Item{CollectionName: u.Collection}

	matchedData := utils.MatchStructureWithSchema(data, u.Schema)

	item.InsertOne(matchedData)
}

// GetClient Returns the client configured with the timeout inverval
func (u *Updater) getClient() (client *http.Client) {
	timeout := time.Duration(u.Timeout) * time.Second

	client = &http.Client{
		Timeout: timeout,
	}

	return
}

// Update Updates the current updater data
func (u *Updater) Update(data map[string]interface{}) {
	// Stop the fetching process
	u.Stop()

	// Check possible values
	if data["interval"] != nil {
		u.Interval = data["interval"].(int)
	}

	method := data["method"]
	if method != "" && (method == "GET" || method == "POST") {
		u.Method = method.(string)
	}

	if data["source"] != "" {
		u.Source = data["source"].(string)
	}

	if data["schema"] != nil {
		u.Schema = data["schema"].(map[string]interface{})
	}

	if data["requestBody"] != nil {
		u.RequestBody = data["requestBody"].(map[string]interface{})
	}

	if data["timeout"] != nil {
		u.Timeout = data["timeout"].(int)
	}
}

// FetchData Fetches the data from the source
func (u *Updater) FetchData() {
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
			response, responseErr := u.getClient().Do(request)

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

					u.SendUpdate(dat)
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
}

// Run Create the scheduler and start running the background taks
func (u *Updater) Run() {
	log.Add(log.Info, "Running updater ", u.ID, uuid.Nil)

	u.scheduler = gocron.NewScheduler(time.UTC)
	u.scheduler.StartAsync()

	u.scheduler.Every(uint64(u.Interval)).Seconds().Do(u.FetchData)
}

// Stop Clears all the background tasks
func (u *Updater) Stop() {
	if u.scheduler != nil {
		u.scheduler.Clear()
	}
}
