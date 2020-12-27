package core

import (
	"net"
	"sync"

	"github.com/cenkalti/rpc2"
	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/src/db"
	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/log"
	"github.com/harvestcore/HarvestCCode/src/updater"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

var lock = &sync.Mutex{}

// UpdaterMap Maps the updater reference with its collection
type UpdaterMap struct {
	Reference  *updater.Updater
	Collection string
}

// Core Main core of the software
type Core struct {
	ID       uuid.UUID
	Updaters map[uuid.UUID]*UpdaterMap

	// RPC
	client     *rpc2.Client
	connection *net.Conn
}

var core *Core

// GetCore Returns the only instance of Core
func GetCore() *Core {
	if core == nil {
		lock.Lock()
		defer lock.Unlock()

		id := uuid.New()

		connection, err := net.Dial("tcp", ":50125")

		if err != nil {
			log.AddSimple(log.Error, "Could not dial port 50125")
		}

		client := rpc2.NewClient(connection)
		registerFunctions(client)
		go client.Run()

		var r utils.Reply
		client.Call("RegisterComponent", utils.RegisterComponentArgs{ComponentType: "CORE", ID: id}, &r)

		if &r != nil {
			log.AddSimple(log.Error, "Could not register Core component")
		}

		core = &Core{
			ID: id,

			// RPC
			client:     client,
			connection: &connection,
		}
	}

	return core
}

// CreateUpdater Creates a new updater and stores a pointer to it.
func (c *Core) CreateUpdater(collection string, schema map[string]interface{}, interval int, source string, method string, requestBody map[string]interface{}, timeout int) uuid.UUID {
	for _, value := range c.Updaters {
		if value.Collection == collection {
			log.AddSimple(log.Error, "Database already in use.")
			log.AddSimple(log.Error, "Updater not created.")
			return uuid.Nil
		}
	}

	var updater = updater.NewUpdater(schema, interval, source, method, requestBody, timeout)

	if updater != nil {
		log.AddSimple(log.Info, "Updater created with ID "+updater.ID.String())
		c.Updaters[updater.ID] = &UpdaterMap{
			Reference:  updater,
			Collection: collection,
		}

		// Start fetching data
		updater.Run()

		return updater.ID
	}

	log.AddSimple(log.Error, "Updater not created.")
	return uuid.Nil
}

// // UpdateUpdater Updates an existing updater.
// func (c *Core) UpdateUpdater(updater uuid.UUID, data *updater.Updater) bool {
// 	var u = c.Updaters[updater]

// 	if u == nil {
// 		log.AddSimple(log.Error, "Updater "+updater.String()+" does not exist.")

// 		return false
// 	}

// 	u.Reference.Update(data)

// 	log.AddSimple(log.Info, "Updater "+updater.String()+" updated.")
// 	return true
// }

// StopUpdater Stops an existing updater and removes it.
func (c *Core) StopUpdater(updater uuid.UUID) bool {
	var u = c.Updaters[updater]

	if u == nil {
		log.AddSimple(log.Error, "Updater "+updater.String()+" does not exist.")

		return false
	}

	// Stop the fetching process
	u.Reference.Stop()

	// Set pointer to null. (GC will free this memory)
	u.Reference = nil

	// Remove entry in the map
	delete(c.Updaters, updater)

	log.AddSimple(log.Info, "Updater "+updater.String()+" removed.")
	return true
}

// StoreData Stores the given data in the given collection
func (c *Core) StoreData(from uuid.UUID, data map[string]interface{}) {
	item := &db.Item{CollectionName: c.Updaters[from].Collection}
	item.InsertOne(data)
}

func (c *Core) FetchData() {

}

func (c *Core) SendEvent() {

}

func registerFunctions(client *rpc2.Client) {
	client.Handle("HandleCoreEvent", func(client *rpc2.Client, e event.Event, reply *utils.Reply) error {
		switch utils.EventType(e.Type) {
		case utils.CreateUpdater:
			// GetCore().CreateUpdater()
		case utils.RemoveUpdater:
			// GetCore().StopUpdater()
		case utils.StoreData:
			GetCore().StoreData(e.From, e.Data)
		}
		return nil
	})
}
