package log_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/HarvestCCode/log"
)

func TestLogCreation(t *testing.T) {
	var niluuid uuid.UUID
	var msg = "yeet"

	var a = uuid.New()
	var b = uuid.New()

	var l = log.New(a, b, msg)
	assert.NotNil(t, l, "Log creation return nil with all parameters set")
	assert.Equal(t, l.From, a, "Mismatch 'from' UUID")
	assert.Equal(t, l.To, b, "Mismatch 'to' UUID")
	assert.Equal(t, l.Message, msg, "Mismatch message")
	assert.NotNil(t, l.ID, "Created log does not have ID")
	assert.NotNil(t, l.Datetime, "Created log does not have Datetime")

	l = log.New(niluuid, uuid.New(), msg)
	assert.Nil(t, l, "Log creation did not return nil when 'from' is a null UUID")

	l = log.New(uuid.New(), niluuid, msg)
	assert.Nil(t, l, "Log creation did not return nil when 'to' is a null UUID")

	l = log.New(niluuid, niluuid, msg)
	assert.Nil(t, l, "Log creation did not return nil when both UUID are null")
}
