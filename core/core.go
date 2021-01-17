package core

import (
	"sync"

	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/log"
	"github.com/harvestcore/HarvestCCode/updater"
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
}

var core *Core

// GetCore Returns the only instance of Core
func GetCore() *Core {
	if core == nil {
		lock.Lock()
		defer lock.Unlock()

		id := uuid.New()

		core = &Core{
			ID:       id,
			Updaters: make(map[uuid.UUID]*UpdaterMap),
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
		data["collection"].(string),
	)

	if updater != nil {
		c.Updaters[updater.ID] = &UpdaterMap{
			Reference:  updater,
			Collection: data["collection"].(string),
		}

		log.AddSimple(log.Info, "Updater created with ID "+updater.ID.String())
		return updater.ID
	}

	log.AddSimple(log.Error, "Updater not created.")
	return uuid.Nil
}

// UpdateUpdater Updates an updater.
func (c *Core) UpdateUpdater(updater uuid.UUID, data map[string]interface{}) {
	u := c.Updaters[updater].Reference
	u.Update(data)

	if data["collection"] != "" && data["collection"] != nil && data["collection"] != c.Updaters[updater].Collection {
		c.Updaters[updater].Collection = data["collection"].(string)
	}
}

// StartUpdater Starts an existing updater.
func (c *Core) StartUpdater(updater uuid.UUID) bool {
	var u = c.Updaters[updater]

	if u == nil {
		log.AddSimple(log.Error, "Updater "+updater.String()+" does not exist.")

		return false
	}

	// Stop the fetching process
	u.Reference.Run()

	log.AddSimple(log.Info, "Updater "+updater.String()+" started.")
	return true
}

// RemoveUpdater Removes an existing updater.
func (c *Core) RemoveUpdater(updater uuid.UUID) bool {
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

// StopUpdater Stops an existing updater.
func (c *Core) StopUpdater(updater uuid.UUID) bool {
	var u = c.Updaters[updater]

	if u == nil {
		log.AddSimple(log.Error, "Updater "+updater.String()+" does not exist.")

		return false
	}

	// Stop the fetching process
	u.Reference.Stop()

	log.AddSimple(log.Info, "Updater "+updater.String()+" stopped.")
	return true
}
