package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/HarvestCCode/src/config"
)

func TestConfigManager(t *testing.T) {
	var cm = config.GetManager()
	assert.NotNil(t, cm, "Config manager creation returned nil")

	assert.NotEqual(t, cm.GetVariable(config.HCC_MONGO_PORT), "", "Mongo port is not empty")
	assert.NotEqual(t, cm.GetVariable(config.HCC_MONGO_HOST), "", "Mongo host is not empty")
	assert.NotEqual(t, cm.GetVariable(config.HCC_MONGO_DATABASE), "", "Mongo database is not empty")
	assert.NotEqual(t, cm.GetVariable(config.HCC_MONGO_URI), "", "Mongo Uri is not empty")
	assert.NotEqual(t, cm.GetVariable(config.HCC_LOG_FILE), "", "Log file is not empty")
	assert.NotEqual(t, cm.GetVariable(config.HCC_ETCD3_HOST), "", "Etcd3 host is not empty")

	var cm2 = config.GetManager()
	assert.Equal(t, cm, cm2, "A 'second' instance is not the same")
}
