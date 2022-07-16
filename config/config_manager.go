package config

import (
	"os"
	"sync"

	"github.com/harvestcore/upgote/types"
)

const (
	// Variable names.
	MONGO_DATABASE   string = "MONGO_DATABASE"
	MONGO_URI        string = "MONGO_URI"
	LOG_FILE         string = "LOG_FILE"
	HTTP_SERVER_PORT string = "HTTP_SERVER_PORT"
	UPGOTE_VERSION   string = "UPGOTE_VERSION"
)

var lock = &sync.Mutex{}

// Manager Encapsulates all the config variables needed.
type Manager struct {
	VariablePool types.Dict
}

var manager *Manager

// GetManager Returns the Manager instance.
func GetManager() *Manager {
	if manager == nil {
		lock.Lock()
		defer lock.Unlock()

		manager = &Manager{
			VariablePool: make(types.Dict),
		}

		manager.setDefaultVariables()
	}

	return manager
}

// getFromEnv Get a variable from the environment.
func getFromEnv(key string, fallback types.Object) types.Object {
	value, exists := os.LookupEnv(key)
	
	if !exists {
        return fallback
    }

    return value
}

// setDefaultVariables Set the default variables.
func (manager *Manager) setDefaultVariables() {
	// Upgote version.
	Set(UPGOTE_VERSION, "0.2.0")

	Set(MONGO_DATABASE, getFromEnv(MONGO_DATABASE, "upgote"))
	Set(MONGO_URI, getFromEnv(MONGO_URI, "mongodb://localhost:27017"))
	Set(LOG_FILE, getFromEnv(LOG_FILE, "/var/log/upgote.log"))
	Set(HTTP_SERVER_PORT, getFromEnv(HTTP_SERVER_PORT, "8080"))
}

// Get Returns the requested variable.
func Get(variable string) types.Object {
	var manager = GetManager()
	var output = manager.VariablePool[string(variable)]

	if output == "" {
		// Try to get the variable from the environment.
		output = getFromEnv(variable, nil)
	}

	return output
}

// Set Set a variable.
func Set(variable string, value types.Object) {
	var manager = GetManager()
	if variable != "" && value != nil {
		manager.VariablePool[variable] = value
	}
}
