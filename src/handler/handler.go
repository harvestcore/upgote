package handler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

type Handler struct {
	EventQueue []event.Event
	Timeout    int
	Scheduler  *gocron.Scheduler
	Lock       bool
}

// NewHandler Creates a new Handler
func NewHandler(timeout int) *Handler {
	var null int

	if timeout == null {
		return nil
	}

	h := Handler{
		Timeout:    timeout,
		Lock:       false,
		EventQueue: make([]event.Event, 0),
		Scheduler:  nil,
	}

	h.StartHandlingEvents()
	return &h
}

// Run Run the handling if the lock is available and the queue has events to process
func (h *Handler) Run() {
	if !h.Lock && len(h.EventQueue) > 0 {
		h.HandleEvents()
	}
}

// StartHandlingEvents Start the process of handling events, aka run the scheduler
func (h *Handler) StartHandlingEvents() {
	h.Scheduler = gocron.NewScheduler(time.UTC)

	h.Scheduler.StartAsync()
	h.Scheduler.Every(uint64(h.Timeout / 2)).Seconds().Do(h.Run)
}

// StopHandlingEvents Stop the handling process
func (h *Handler) StopHandlingEvents() {
	if h.Scheduler != nil {
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

func (h *Handler) QueueEvent(e event.Event) {
	h.EventQueue = append(h.EventQueue, e)
}

func (h *Handler) SendEvent() {

}
