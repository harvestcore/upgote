package log_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/harvestcore/upgote/log"
)

func TestLogCreation(t *testing.T) {
	var msg = "yeet"

	var a = uuid.New()
	var b = uuid.New()

	var l = log.NewLog(log.Error, msg, a, b)
	assert.NotNil(t, l, "Log creation return nil with all parameters set")
	assert.Equal(t, l.From, a, "Mismatch 'from' UUID")
	assert.Equal(t, l.To, b, "Mismatch 'to' UUID")
	assert.Equal(t, l.Message, msg, "Mismatch message")
	assert.NotNil(t, l.ID, "Created log does not have ID")
	assert.NotNil(t, l.Datetime, "Created log does not have Datetime")

	l = log.NewLog(log.Error, msg, uuid.Nil, uuid.Nil)
	assert.NotNil(t, l, "Log creation return nil with all parameters set")
	assert.Equal(t, l.From, uuid.Nil, "Mismatch 'from' UUID")
	assert.Equal(t, l.To, uuid.Nil, "Mismatch 'to' UUID")
	assert.Equal(t, l.Message, msg, "Mismatch message")
	assert.NotNil(t, l.ID, "Created log does not have ID")
	assert.NotNil(t, l.Datetime, "Created log does not have Datetime")
}
