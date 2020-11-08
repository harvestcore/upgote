package log

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	Datetime time.Time
	From     uuid.UUID
	To       uuid.UUID
	Message  string
	ID       uuid.UUID
}

// New Creates a new Log
func New(from uuid.UUID, to uuid.UUID, message string) *Log {
	var null uuid.UUID

	if from == null || to == null {
		return nil
	}

	return &Log{
		Datetime: time.Now(),
		From:     from,
		To:       to,
		Message:  message,
		ID:       uuid.New(),
	}
}
