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
	HCC_MONGO_PORT     Variable = "HCC_MONGO_PORT"
	HCC_MONGO_HOST     Variable = "HCC_MONGO_HOST"
	HCC_MONGO_DATABASE Variable = "HCC_MONGO_DATABASE"
	HCC_MONGO_URI      Variable = "HCC_MONGO_URI"
)

var lock = &sync.Mutex{}

// Manager Encapsulated all the config variables needed
type Manager struct {
	MongoPort  string
	MongoHost  string
	MongoURI   string
	Database   string
	Context    context.Context
	CancelFunc context.CancelFunc
	Client     *clientv3.Client
	KV         clientv3.KV
}

var manager *Manager

// GetManager Returns the Manager instance
func GetManager() *Manager {
	if manager == nil {
		lock.Lock()
		defer lock.Unlock()

		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, _ := clientv3.New(clientv3.Config{
			DialTimeout: 2 * time.Second,
			Endpoints:   []string{"127.0.0.1:2379"},
		})

		var port = os.Getenv("HCC_MONGO_PORT")
		if port == "" {
			port = "27017"
		}

		var host = os.Getenv("HCC_MONGO_HOST")
		if host == "" {
			host = "localhost"
		}

		var database = os.Getenv("HCC_MONGO_DATABASE")
		if database == "" {
			database = "harvestccode"
		}

		var mongoURI = os.Getenv("HCC_MONGO_URI")
		if mongoURI == "" {
			mongoURI = "mongodb://" + host + ":" + port
		}

		manager = &Manager{
			MongoPort:  port,
			MongoHost:  host,
			MongoURI:   mongoURI,
			Database:   database,
			Context:    context,
			CancelFunc: cancel,
			Client:     client,
			KV:         clientv3.NewKV(client),
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
	var _var string
	data, _ := manager.KV.Get(manager.Context, string(variable))
	for _, ev := range data.Kvs {
		_var = string(ev.Value)
	}

	return _var
}
