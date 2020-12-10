package utils

type EventType string

const (
	Store      EventType = "store"
	Fetch      EventType = "fetch"
	Update     EventType = "update"
	APIsend    EventType = "APIsend"
	APIreceive EventType = "APIreceive"
)
