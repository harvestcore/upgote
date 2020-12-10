package config

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type Variable string

const (
	// Variable names
	HCC_MONGO_PORT     Variable = "HCC_MONGO_PORT"
	HCC_MONGO_HOST     Variable = "HCC_MONGO_HOST"
	HCC_MONGO_DATABASE Variable = "HCC_MONGO_DATABASE"
	HCC_MONGO_URI      Variable = "HCC_MONGO_URI"
	HCC_LOG_FILE       Variable = "HCC_LOG_FILE"
	HCC_ETCD3_HOST     Variable = "HCC_ETCD3_HOST"

	// Default variables
	HCC_DEFAULT_MONGO_PORT     Variable = "27017"
	HCC_DEFAULT_MONGO_HOST     Variable = "localhost"
	HCC_DEFAULT_MONGO_DATABASE Variable = "harvestccode"
	HCC_DEFAULT_MONGO_URI      Variable = "mongodb://localhost:27017"
	HCC_DEFAULT_LOG_FILE       Variable = "/var/log/harvestccode/logfile"
	HCC_DEFAULT_ETCD3_HOST     Variable = "127.0.0.1:2379"
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
		return string(HCC_DEFAULT_LOG_FILE)
	case HCC_ETCD3_HOST:
		return string(HCC_DEFAULT_ETCD3_HOST)
	}

	return ""
}

var lock = &sync.Mutex{}

// Manager Encapsulates all the config variables needed
type Manager struct {
	VariablePool map[string]string
	Context      context.Context
	CancelFunc   context.CancelFunc
	Client       *clientv3.Client
	KV           clientv3.KV
}

var manager *Manager

// GetManager Returns the Manager instance
func GetManager() *Manager {
	if manager == nil {
		lock.Lock()
		defer lock.Unlock()

		var etcd3Host = os.Getenv(string(HCC_MONGO_PORT))

		if etcd3Host == "" {
			// Default one
			etcd3Host = string(HCC_DEFAULT_ETCD3_HOST)
		}

		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, _ := clientv3.New(clientv3.Config{
			DialTimeout: 2 * time.Second,
			Endpoints:   []string{etcd3Host},
		})

		manager = &Manager{
			VariablePool: make(map[string]string),
			Context:      context,
			CancelFunc:   cancel,
			Client:       client,
			KV:           clientv3.NewKV(client),
		}
	}

	return manager
}

// Close Closes the etcd manager
func (manager *Manager) Close() {
	manager.CancelFunc()
}

// GetVariable Returns the requested variable.
func (manager *Manager) GetVariable(variable Variable) string {
	var output = manager.VariablePool[string(variable)]

	if output == "" {
		data, err := manager.KV.Get(manager.Context, string(variable))

		if err == nil {
			for _, ev := range data.Kvs {
				output = string(ev.Value)
			}
		} else {
			output = os.Getenv(string(variable))
		}

		if output == "" {
			output = GetDefault(variable)
		}

		if output != "" {
			manager.VariablePool[string(variable)] = string(variable)
		}
	}

	return output
}
