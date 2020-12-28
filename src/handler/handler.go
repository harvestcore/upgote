package handler

import (
	"net"
	"sync"
	"time"

	"github.com/cenkalti/rpc2"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"

	"github.com/harvestcore/HarvestCCode/src/event"
	"github.com/harvestcore/HarvestCCode/src/log"
	"github.com/harvestcore/HarvestCCode/src/utils"
)

// API API type
type API int

var lock = &sync.Mutex{}

// Handler Event handler
type Handler struct {
	EventQueue []event.Event
	Timeout    int
	Scheduler  *gocron.Scheduler
	Lock       bool

	// RPC
	server   *rpc2.Server
	listener *net.Listener

	updaters   map[uuid.UUID]*rpc2.Client
	apiClient  *rpc2.Client
	coreClient *rpc2.Client
}

var handler *Handler

// GetHandler Returns the only instance of the Handler
func GetHandler() *Handler {
	if handler == nil {
		lock.Lock()
		defer lock.Unlock()

		server := rpc2.NewServer()
		registerFunctions(server)

		listener, err := net.Listen("tcp", ":50125")

		if err != nil {
			log.AddSimple(log.Error, "Error listening on port 50125")
		}

		go server.Accept(listener)

		handler = &Handler{
			Timeout:    1,
			Lock:       false,
			EventQueue: make([]event.Event, 0),
			Scheduler:  gocron.NewScheduler(time.UTC),

			// RPC2
			server:   server,
			listener: &listener,
			updaters: make(map[uuid.UUID]*rpc2.Client, 0),
		}

		log.AddSimple(log.Info, "Created handler.")
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
			var reply utils.Reply
			evtType := utils.EventType(event.Type)

			switch evtType {
			case utils.CreateUpdater:
			case utils.UpdateUpdater:
			case utils.RemoveUpdater:
			case utils.StoreData:
				if h.coreClient != nil {
					h.coreClient.Call("HandleCoreEvent", event, &reply)
				} else {
					log.NewLog(log.Error, "Core client does not exists.", event.From, event.To)
				}
			case utils.API:
				if h.coreClient != nil {
					h.apiClient.Call("HandleAPIEvent", event, &reply)
				} else {
					log.NewLog(log.Error, "API client does not exists.", event.From, event.To)
				}
			case utils.Updater:
				updaterClient := h.updaters[event.From]
				if updaterClient != nil {
					updaterClient.Call("HandleUpdaterEvent", event, &reply)
				} else {
					log.NewLog(log.Error, "Updater client does not exists.", event.From, event.To)
				}

			}
		}
	}

	h.Lock = false
}

// ClearEventQueue Clears all the queued events
func (h *Handler) ClearEventQueue() {
	h.EventQueue = h.EventQueue[:0]
}

// EVENTS

// registerFunctions Register the functions that will be available for the other processes.
func registerFunctions(server *rpc2.Server) {

	// Add a new event to the events pool
	server.Handle("QueueEvent", func(client *rpc2.Client, e *event.Event, reply *utils.Reply) error {
		handler = GetHandler()
		handler.EventQueue = append(handler.EventQueue, *e)

		return nil
	})

	// Used to register all the components (or processes) that will be managed by this handler.
	// There are only three types: Updater, Core and API
	server.Handle("RegisterComponent", func(client *rpc2.Client, args *utils.RegisterComponentArgs, reply *utils.Reply) error {
		handler = GetHandler()

		if args.ComponentType == "UPDATER" {
			handler.updaters[args.ID] = client
		} else if args.ComponentType == "CORE" {
			handler.coreClient = client
		} else if args.ComponentType == "API" {
			handler.apiClient = client
		}

		log.AddSimple(log.Info, "Registered "+args.ComponentType+" component with ID "+args.ID.String())
		return nil
	})

	// Used to unregister the components (or processes) that are managed by this handler.
	// There are only three types: Updater, Core and API
	server.Handle("UnregisterComponent", func(client *rpc2.Client, args *utils.RegisterComponentArgs, reply *utils.Reply) error {
		handler = GetHandler()

		if args.ComponentType == "UPDATER" {
			delete(handler.updaters, args.ID)
		} else if args.ComponentType == "CORE" {
			handler.coreClient = nil
		} else if args.ComponentType == "API" {
			handler.apiClient = nil
		}

		log.AddSimple(log.Info, "Unregistered "+args.ComponentType+" component with ID "+args.ID.String())
		return nil
	})
}
