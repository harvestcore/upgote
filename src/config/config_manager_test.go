package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/HarvestCCode/src/config"
)

func TestConfigManager(t *testing.T) {
	var cm = config.GetManager()
	assert.NotNil(t, cm, "Config manager creation returned nil")

	assert.NotEqual(t, cm.Database, "", "Database is empty")
	assert.NotEqual(t, cm.MongoPort, "", "MongoPort is empty")
	assert.NotEqual(t, cm.MongoHost, "", "MongoHost is empty")
	assert.NotEqual(t, cm.MongoURI, "", "MongoURI is empty")

	var cm2 = config.GetManager()
	assert.Equal(t, cm, cm2, "A 'second' instance is not the same")
}
