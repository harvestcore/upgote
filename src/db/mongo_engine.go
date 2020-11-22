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
	cancel   context.CancelFunc
	ctx      context.Context
}

var engine *MongoEngine

// GetEngine Returns the only instance of MongoEngine
func GetEngine() *MongoEngine {
	if engine == nil {
		lock.Lock()
		defer lock.Unlock()

		// Config variables
		var mongoURI = config.GetManager().MongoURI
		var database = config.GetManager().Database

		// Mongo client instantation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

		if err == nil {
			engine = &MongoEngine{
				Client:   client,
				Database: database,
				cancel:   cancel,
				ctx:      ctx,
			}
		} else {
			defer func() {
				cancel()
				client.Disconnect(ctx)
			}()
		}
	}

	return engine
}

// CloseConnection Closes the client connectino
func CloseConnection() {
	var engine = GetEngine()
	engine.cancel()
	engine.Client.Disconnect(engine.ctx)
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
