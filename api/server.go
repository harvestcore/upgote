package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	handlers "github.com/harvestcore/upgote/api/handlers"
	middlewares "github.com/harvestcore/upgote/api/middlewares"
	"github.com/harvestcore/upgote/config"
	"github.com/harvestcore/upgote/log"
)

var lock = &sync.Mutex{}

// Server HTTP server of the software.
type Server struct {
	Server *http.Server
}

var server *Server

// GetServer Returns the only instance of the HTTP server.
func GetServer() *Server {
	if server == nil {
		lock.Lock()
		defer lock.Unlock()

		router := mux.NewRouter().PathPrefix("/api").Subrouter()

		registerHandlers(router)
		registerMiddlewares(router)

		server = &Server{
			Server: &http.Server{
				Handler: router,
				Addr:    ":" + config.Get(config.HTTP_SERVER_PORT).(string),

				// Read and write timeouts to avoid the server hang.
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			},
		}
	}

	return server
}

// Start Starts listening and serving requests.
func (s *Server) Start() {
	log.AddSimple(log.Info, "HTTP Server started, running on address "+s.Server.Addr)

	if err := s.Server.ListenAndServe(); err != nil {
		log.AddSimple(log.Error, "Error running API server listening on port "+s.Server.Addr)
	}
}

// registerHandlers Registers all the server handlers.
func registerHandlers(router *mux.Router) {
	handlers.Healthcheck(router)
	handlers.Status(router)
	handlers.Log(router)
	handlers.Data(router)
	handlers.Updater(router)
}

// registerMiddlewares Registers all the server middlewares.
func registerMiddlewares(router *mux.Router) {
	router.Use(middlewares.LoggingMiddleware)
}
