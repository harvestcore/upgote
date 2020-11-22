package config

import (
	"os"
	"sync"
)

var lock = &sync.Mutex{}

// Manager Encapsulated all the config variables needed
type Manager struct {
	MongoPort string
	MongoHost string
	MongoURI  string
	Database  string
}

var manager *Manager

// GetManager Returns the Manager instance
func GetManager() *Manager {
	if manager == nil {
		lock.Lock()
		defer lock.Unlock()

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
			MongoPort: port,
			MongoHost: host,
			MongoURI:  mongoURI,
			Database:  database,
		}
	}

	return manager
}
