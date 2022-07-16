package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/upgote/api"
)

func TestServerCreation(t *testing.T) {
	var s = api.GetServer()
	assert.NotNil(t, s, "Server creation failed")
}
