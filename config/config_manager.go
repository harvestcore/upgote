package config

import (
	"os"
	"sync"
)

type Variable string

const (
	// Variable names
	MONGO_DATABASE   Variable = "MONGO_DATABASE"
	MONGO_URI        Variable = "MONGO_URI"
	LOG_FILE         Variable = "LOG_FILE"
	HTTP_SERVER_PORT Variable = "HTTP_SERVER_PORT"

	// Default variables
	DEFAULT_MONGO_DATABASE   Variable = "upgote"
	DEFAULT_MONGO_URI        Variable = "mongodb://127.0.0.1:27017"
	DEFAULT_LOG_FILE         Variable = "/upgote.log"
	DEFAULT_HTTP_SERVER_PORT Variable = "80"
)

// GetDefault Returns the default value of a variable
func GetDefault(variable Variable) string {
	switch variable {
	case MONGO_URI:
		return string(DEFAULT_MONGO_URI)
	case MONGO_DATABASE:
		return string(DEFAULT_MONGO_DATABASE)
	case LOG_FILE:
		return os.Getenv("HOME") + string(DEFAULT_LOG_FILE)
	case HTTP_SERVER_PORT:
		return string(DEFAULT_HTTP_SERVER_PORT)
	}

	return ""
}

var lock = &sync.Mutex{}

// Manager Encapsulates all the config variables needed
type Manager struct {
	VariablePool map[string]string
}

var manager *Manager

// GetManager Returns the Manager instance
func GetManager() *Manager {
	if manager == nil {
		lock.Lock()
		defer lock.Unlock()

		manager = &Manager{
			VariablePool: make(map[string]string),
		}
	}

	return manager
}


// GetVariable Returns the requested variable.
func (manager *Manager) GetVariable(variable Variable) string {
	var output = manager.VariablePool[string(variable)]

	// Try to get the variable from remote.
	if output == "" {
		// Try to get the variable from the environment.
		output = os.Getenv(string(variable))

		// Default value.
		if output == "" {
			output = GetDefault(variable)
		} else {
			manager.VariablePool[string(variable)] = string(output)
		}
	}

	return output
}
