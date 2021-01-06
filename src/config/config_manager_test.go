package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/HarvestCCode/src/config"
)

func TestConfigManager(t *testing.T) {
	var cm = config.GetManager()
	assert.NotNil(t, cm, "Config manager creation returned nil")

	var cm2 = config.GetManager()
	assert.Equal(t, cm, cm2, "A 'second' instance is not the same")
}
