package log

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/harvestcore/HarvestCCode/src/config"
	"github.com/harvestcore/HarvestCCode/src/db"
)

type Connotation string

const (
	Info    Connotation = "INFO"
	Warning Connotation = "WARNING"
	Error   Connotation = "ERROR"
)

// Log Encapsulates a log message
type Log struct {
	Connotation Connotation
	Datetime    time.Time
	From        uuid.UUID
	To          uuid.UUID
	Message     string
	ID          uuid.UUID
}

// NewLog Creates a new Log
func NewLog(connotation Connotation, from uuid.UUID, to uuid.UUID, message string) *Log {
	var null uuid.UUID

	if from == null || to == null {
		return nil
	}

	return &Log{
		Connotation: connotation,
		Datetime:    time.Now(),
		From:        from,
		To:          to,
		Message:     message,
		ID:          uuid.New(),
	}
}

func (l *Log) serialize() map[string]interface{} {
	return map[string]interface{}{
		"connotation": string(l.Connotation),
		"datetime":    l.Datetime.String(),
		"from":        l.From.String(),
		"to":          l.To.String(),
		"message":     l.Message,
		"id":          l.ID.String(),
	}
}

// Logger Encapsulates the logging system
type Logger struct {
	LogFile *os.File
	Item    *db.Item
}

// Main logger
var logger *Logger

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

// getLogger Returns the Logger instance
func getLogger() *Logger {
	if logger == nil {
		logFilePath := config.GetManager().GetVariable(config.HCC_LOG_FILE)
		file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err == nil {
			logger = &Logger{
				LogFile: file,
				// Database connection
				Item: &db.Item{
					CollectionName: "log",
				},
			}

			// Set logger properties
			infoLogger = log.New(file, "[INFO] - ", log.LstdFlags|log.Lshortfile)
			warningLogger = log.New(file, "[WARNING] - ", log.LstdFlags|log.Lshortfile)
			errorLogger = log.New(file, "[ERROR] - ", log.LstdFlags|log.Lshortfile)
		}
	}

	return logger
}

// Stop Stops the logger
// Closes the log file
func (logger *Logger) Stop() {
	logger.LogFile.Close()
}

// Add Adds a new log message
func Add(connotation Connotation, message string, from uuid.UUID, to uuid.UUID) {
	var logger = getLogger()
	var _log = NewLog(connotation, from, to, message)

	if _log != nil {
		var _message = message + "__from__" + from.String() + "__to__" + to.String()

		// Add log message to local text file
		switch connotation {
		case Info:
			infoLogger.Println(_message)
		case Warning:
			warningLogger.Println(_message)
		case Error:
			errorLogger.Println(_message)
		}

		// Add log message to database
		logger.Item.InsertOne(_log.serialize())
	}
}
