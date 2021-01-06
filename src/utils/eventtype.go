package utils

type EventType string

const (
	CreateUpdater EventType = "CreateUpdater"
	RemoveUpdater EventType = "RemoveUpdater"
	UpdateUpdater EventType = "UpdateUpdater"
	StartUpdater  EventType = "StartUpdater"
	StoreData     EventType = "StoreData"
	Updater       EventType = "Updater"
	API           EventType = "API"
)
