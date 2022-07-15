package db_test

import (
	"testing"

	"github.com/harvestcore/upgote/db"
	"github.com/harvestcore/upgote/utils"
	"github.com/stretchr/testify/assert"
)

func TestMongoEngine(t *testing.T) {
	if !utils.RunningInDocker() {
		engine := db.GetEngine()

		assert.NotEqual(t, engine, nil, "MongoEngine is nil")
		assert.NotEqual(t, engine.Client, nil, "The client is not initialized")
		assert.NotEqual(t, engine.Database, "", "The database name is empty")
	}
}
