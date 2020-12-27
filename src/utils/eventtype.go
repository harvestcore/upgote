package utils

type EventType string

const (
	CreateUpdater EventType = "CreateUpdater"
	RemoveUpdater EventType = "RemoveUpdater"
	UpdateUpdater EventType = "UpdateUpdater"
	StoreData     EventType = "StoreData"
	API           EventType = "API"
)
