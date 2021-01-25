package config

import (
	"context"
	"net"
	"os"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Variable string

const (
	// Variable names
	HCC_MONGO_DATABASE   Variable = "HCC_MONGO_DATABASE"
	HCC_MONGO_URI        Variable = "HCC_MONGO_URI"
	HCC_LOG_FILE         Variable = "HCC_LOG_FILE"
	HCC_HTTP_SERVER_PORT Variable = "HCC_HTTP_SERVER_PORT"
	HCC_ETCD3_HOST       Variable = "HCC_ETCD3_HOST"

	// Default variables
	HCC_DEFAULT_MONGO_DATABASE   Variable = "harvestccode"
	HCC_DEFAULT_MONGO_URI        Variable = "mongodb://127.0.0.1:27017"
	HCC_DEFAULT_LOG_FILE         Variable = "/harvestccode.log"
	HCC_DEFAULT_HTTP_SERVER_PORT Variable = "8080"
	HCC_DEFAULT_ETCD3_HOST       Variable = "127.0.0.1:2379"
)

// GetDefault Returns the default value of a variable
func GetDefault(variable Variable) string {
	switch variable {
	case HCC_MONGO_URI:
		return string(HCC_DEFAULT_MONGO_URI)
	case HCC_MONGO_DATABASE:
		return string(HCC_DEFAULT_MONGO_DATABASE)
	case HCC_LOG_FILE:
		return os.Getenv("HOME") + string(HCC_DEFAULT_LOG_FILE)
	case HCC_HTTP_SERVER_PORT:
		return string(HCC_DEFAULT_HTTP_SERVER_PORT)
	}

	return ""
}

var lock = &sync.Mutex{}

// Manager Encapsulates all the config variables needed
type Manager struct {
	VariablePool map[string]string

	// Etcd3
	Context      context.Context
	CancelFunc   context.CancelFunc
	Client       *clientv3.Client
	KV           clientv3.KV
	Etcd3Enabled bool
}

var manager *Manager

// GetManager Returns the Manager instance
func GetManager() *Manager {
	if manager == nil {
		lock.Lock()
		defer lock.Unlock()

		var etcd3Enabled = true
		var ctx context.Context
		var cancel context.CancelFunc
		var client *clientv3.Client
		var kv clientv3.KV
		var etcd3Host = os.Getenv(string(HCC_ETCD3_HOST))

		if etcd3Host == "" {
			// Default one
			etcd3Host = string(HCC_DEFAULT_ETCD3_HOST)
		}

		_, err := net.DialTimeout("tcp", etcd3Host, 2*time.Second)
		if err != nil {
			etcd3Enabled = false
		}

		if etcd3Enabled {
			ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
			client, _ = clientv3.New(clientv3.Config{
				DialTimeout: 2 * time.Second,
				Endpoints:   []string{etcd3Host},
			})

			kv = clientv3.NewKV(client)
		}

		manager = &Manager{
			VariablePool: make(map[string]string),

			Context:      ctx,
			CancelFunc:   cancel,
			Client:       client,
			KV:           kv,
			Etcd3Enabled: etcd3Enabled,
		}
	}

	return manager
}

// Close Closes the etcd manager.
func (manager *Manager) Close() {
	manager.CancelFunc()
}

// GetFromRemote Gets the variable from ETCD3 server.
func (manager *Manager) GetFromRemote(variable Variable) string {
	var output string

	if manager.Etcd3Enabled {
		data, err := manager.KV.Get(manager.Context, string(variable))

		if err == nil {
			for _, ev := range data.Kvs {
				output = string(ev.Value)
			}
		} else {
			output = os.Getenv(string(variable))
		}
	}

	return output
}

// GetVariable Returns the requested variable.
func (manager *Manager) GetVariable(variable Variable) string {
	var output = manager.VariablePool[string(variable)]

	// Try to get the variable from remote.
	if output == "" {
		output = manager.GetFromRemote(variable)

		// Try to get the variable from the environment.
		if output == "" {
			output = os.Getenv(string(variable))

			// Default value.
			if output == "" {
				output = GetDefault(variable)
			}
		}

		if output != "" {
			manager.VariablePool[string(variable)] = string(output)
		}
	}

	return output
}
