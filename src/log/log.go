package log

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	datetime time.Time
	from     uuid.UUID
	message  string
	id       uuid.UUID
}
