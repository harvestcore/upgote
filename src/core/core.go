package core

import (
	"sync"

	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/src/log"
	"github.com/harvestcore/HarvestCCode/src/updater"
)

var lock = &sync.Mutex{}

// UpdaterMap Maps the updater reference with its collection
type UpdaterMap struct {
	Reference  *updater.Updater
	Collection string
}

// Core Main core of the software
type Core struct {
	Updaters map[uuid.UUID]*UpdaterMap
}

var core *Core

// GetCore Returns the only instance of Core
func GetCore() *Core {
	if core == nil {
		lock.Lock()
		defer lock.Unlock()

		core = &Core{
			Updaters: make(map[uuid.UUID]*UpdaterMap, 0),
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

func (c *Core) UpdateData() {

}

func (c *Core) StoreData() {

}

func (c *Core) FetchData() {

}

func (c *Core) HandleEvent() {

}

func (c *Core) SendEvent() {

}
