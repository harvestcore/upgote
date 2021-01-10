package core

import (
	"net"
	"sync"

	"github.com/cenkalti/rpc2"
	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/src/config"
	"github.com/harvestcore/HarvestCCode/src/db"
	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/handler"
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
		var client *rpc2.Client

		// Wake up handler
		handler.GetHandler()

		port := config.GetManager().GetVariable(config.HCC_RPC_PORT)
		connection, err := net.Dial("tcp", ":"+port)

		if err != nil {
			log.AddSimple(log.Error, "Could not dial port "+port)
		} else {
			client := rpc2.NewClient(connection)
			registerFunctions(client)
			go client.Run()

			var r utils.Reply
			client.Call("RegisterComponent", utils.RegisterComponentArgs{ComponentType: "CORE", ID: id}, &r)

			if &r != nil {
				log.AddSimple(log.Error, "Could not register Core component")
			}
		}

		core = &Core{
			ID:       id,
			Updaters: make(map[uuid.UUID]*UpdaterMap),

			// RPC
			client:     client,
			connection: &connection,
		}
	}

	return core
}

// CreateUpdater Creates a new updater and stores a pointer to it.
func (c *Core) CreateUpdater(data map[string]interface{}) uuid.UUID {
	for _, value := range c.Updaters {
		if value.Collection == data["collection"].(string) {
			log.AddSimple(log.Error, "Database already in use.")
			log.AddSimple(log.Error, "Updater not created.")
			return uuid.Nil
		}
	}

	var updater = updater.NewUpdater(
		data["schema"].(map[string]interface{}),
		data["interval"].(int),
		data["source"].(string),
		data["method"].(string),
		data["requestBody"].(map[string]interface{}),
		data["timeout"].(int),
	)

	if updater != nil {
		log.AddSimple(log.Info, "Updater created with ID "+updater.ID.String())
		c.Updaters[updater.ID] = &UpdaterMap{
			Reference:  updater,
			Collection: data["collection"].(string),
		}

		return updater.ID
	}

	log.AddSimple(log.Error, "Updater not created.")
	return uuid.Nil
}

// UpdateUpdater Updates an updater.
func (c *Core) UpdateUpdater(to uuid.UUID, data map[string]interface{}) {
	var reply utils.Reply
	d := make(map[string]interface{})
	d["data"] = data
	d["reference"] = c.Updaters[to].Reference

	e := event.NewEvent(uuid.Nil, to, utils.UpdateUpdater, "", d)
	c.client.Call("QueueEvent", e, &reply)
}

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

func registerFunctions(client *rpc2.Client) {
	client.Handle("HandleCoreEvent", func(client *rpc2.Client, e *event.Event, reply *utils.Reply) error {
		switch utils.EventType(e.Type) {
		case utils.CreateUpdater:
			GetCore().CreateUpdater(e.Data)
		case utils.UpdateUpdater:
			GetCore().UpdateUpdater(e.To, e.Data)
		case utils.RemoveUpdater:
			GetCore().StopUpdater(e.Data["updater"].(uuid.UUID))
		case utils.StoreData:
			GetCore().StoreData(e.From, e.Data)
		}

		return nil
	})
}
