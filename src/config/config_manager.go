package config

import (
	"context"
	"os"
	"sync"
)

type Variable string

const (
	// Variable names
	HCC_MONGO_PORT       Variable = "HCC_MONGO_PORT"
	HCC_MONGO_HOST       Variable = "HCC_MONGO_HOST"
	HCC_MONGO_DATABASE   Variable = "HCC_MONGO_DATABASE"
	HCC_MONGO_URI        Variable = "HCC_MONGO_URI"
	HCC_LOG_FILE         Variable = "HCC_LOG_FILE"
	HCC_ETCD3_HOST       Variable = "HCC_ETCD3_HOST"
	HCC_RPC_PORT         Variable = "HCC_RPC_PORT"
	HCC_HTTP_SERVER_PORT Variable = "HCC_HTTP_SERVER_PORT"

	// Default variables
	HCC_DEFAULT_MONGO_PORT       Variable = "27017"
	HCC_DEFAULT_MONGO_HOST       Variable = "localhost"
	HCC_DEFAULT_MONGO_DATABASE   Variable = "harvestccode"
	HCC_DEFAULT_MONGO_URI        Variable = "mongodb://127.0.0.1:27017"
	HCC_DEFAULT_LOG_FILE         Variable = "/harvestccode.log"
	HCC_DEFAULT_ETCD3_HOST       Variable = "127.0.0.1:2379"
	HCC_DEFAULT_RPC_PORT         Variable = "50125"
	HCC_DEFAULT_HTTP_SERVER_PORT Variable = "8080"
)

// GetDefault Returns the default value of a variable
func GetDefault(variable Variable) string {
	switch variable {
	case HCC_MONGO_PORT:
		return string(HCC_DEFAULT_MONGO_PORT)
	case HCC_MONGO_HOST:
		return string(HCC_DEFAULT_MONGO_HOST)
	case HCC_MONGO_DATABASE:
		return string(HCC_DEFAULT_MONGO_DATABASE)
	case HCC_MONGO_URI:
		return string(HCC_DEFAULT_MONGO_URI)
	case HCC_LOG_FILE:
		return os.Getenv("HOME") + string(HCC_DEFAULT_LOG_FILE)
	case HCC_ETCD3_HOST:
		return string(HCC_DEFAULT_ETCD3_HOST)
	case HCC_RPC_PORT:
		return string(HCC_DEFAULT_RPC_PORT)
	case HCC_HTTP_SERVER_PORT:
		return string(HCC_DEFAULT_HTTP_SERVER_PORT)
	}

	return ""
}

var lock = &sync.Mutex{}

// Manager Encapsulates all the config variables needed
type Manager struct {
	VariablePool map[string]string
	Context      context.Context
	CancelFunc   context.CancelFunc
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

	if output == "" {
		output = os.Getenv(string(variable))

		if output == "" {
			output = GetDefault(variable)
		}

		if output != "" {
			manager.VariablePool[string(variable)] = string(output)
		}
	}

	return output
}
