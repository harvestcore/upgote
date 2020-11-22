package db

import (
	"context"
	"sync"
	"time"

	"github.com/harvestcore/HarvestCCode/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}

// MongoEngine Encapsulates a Mongo client with its database name
type MongoEngine struct {
	Client   *mongo.Client
	Database string
}

var engine *MongoEngine

// GetEngine Returns the only instance of MongoEngine
func GetEngine() *MongoEngine {
	if engine == nil {
		lock.Lock()
		defer lock.Unlock()

		// Config variables
		var mongoURI = config.GetConfigManager().MongoURI
		var database = config.GetConfigManager().Database

		// Mongo client instantation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()

		engine = &MongoEngine{
			Client:   client,
			Database: database,
		}
	}

	return engine
}

// DB Returns the current active database
func DB() *mongo.Database {
	var engine = GetEngine()
	return engine.Client.Database(engine.Database)
}

// Ctx Returns the context
func Ctx() context.Context {
	return context.TODO()
}
