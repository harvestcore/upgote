package api_test

import (
	"testing"

	"github.com/harvestcore/upgote/api"
	"github.com/stretchr/testify/assert"
)

func TestServerCreation(t *testing.T) {
	var s = api.GetServer()
	assert.NotNil(t, s, "Server creation failed")
}
