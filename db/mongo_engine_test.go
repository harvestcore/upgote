package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/upgote/db"
)

func TestMongoEngine(t *testing.T) {
	engine := db.GetEngine()

	assert.NotEqual(t, engine, nil, "MongoEngine is nil")
	assert.NotEqual(t, engine.Client, nil, "The client is not initialized")
	assert.NotEqual(t, engine.Database, "", "The database name is empty")
}
