package handler

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/log"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

var lock = &sync.Mutex{}

// Handler Event handler
type Handler struct {
	EventQueue []event.Event
	Timeout    int
	Scheduler  *gocron.Scheduler
	Lock       bool
}

var handler *Handler

// GetHandler Returns the only instance of the Handler
func GetHandler() *Handler {
	if handler == nil {
		lock.Lock()
		defer lock.Unlock()

		log.AddSimple(log.Info, "Created handler.")
		handler = &Handler{
			Timeout:    1,
			Lock:       false,
			EventQueue: make([]event.Event, 0),
			Scheduler:  gocron.NewScheduler(time.UTC),
		}
	}

	return handler
}

// Run Run the handling if the lock is available and the queue has events to process
func (h *Handler) Run() {
	if !h.Lock && len(h.EventQueue) > 0 {
		h.HandleEvents()
	}
}

// StartHandlingEvents Start the process of handling events, aka run the scheduler
func (h *Handler) StartHandlingEvents() {
	log.AddSimple(log.Info, "Start handling events.")

	h.Scheduler.StartAsync()
	h.Scheduler.Every(uint64(h.Timeout / 2)).Seconds().Do(h.Run)
}

// StopHandlingEvents Stop the handling process
func (h *Handler) StopHandlingEvents() {
	if h.Scheduler != nil {
		log.AddSimple(log.Info, "Stop handling events.")

		h.Scheduler.Clear()
	}
}

// HandleEvents Handle each type of event
func (h *Handler) HandleEvents() {
	h.Lock = true
	for {
		if len(h.EventQueue) == 0 {
			break
		} else {
			event := h.EventQueue[0]
			h.EventQueue = h.EventQueue[1:]

			evtType := utils.EventType(event.Type)

			switch evtType {
			case utils.Store:
				fmt.Println("Store")
			case utils.Fetch:
				fmt.Println("Fetch")
			case utils.Update:
				fmt.Println("Update")
			case utils.APIsend:
				fmt.Println("APIsend")
			case utils.APIreceive:
				fmt.Println("APIreceive")
			}
		}
	}

	h.Lock = false
}

// ClearEventQueue Clears all the queued events
func (h *Handler) ClearEventQueue() {
	h.EventQueue = h.EventQueue[:0]
}

// QueueEvent Queues a new event
func (h *Handler) QueueEvent(e event.Event) {
	h.EventQueue = append(h.EventQueue, e)
}

func (h *Handler) SendEvent() {

}
