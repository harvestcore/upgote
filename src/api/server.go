package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	handlers "github.com/harvestcore/HarvestCCode/src/api/handlers"
	"github.com/harvestcore/HarvestCCode/src/config"
	"github.com/harvestcore/HarvestCCode/src/log"
)

var lock = &sync.Mutex{}

// Server HTTP server of the software
type Server struct {
	Server *http.Server
}

var server *Server

// GetServer Returns the only instance of the HTTP server
func GetServer() *Server {
	if server == nil {
		lock.Lock()
		defer lock.Unlock()

		router := mux.NewRouter()

		registerHandlers(router)

		server = &Server{
			Server: &http.Server{
				Handler: router,
				Addr:    ":" + config.GetManager().GetVariable(config.HCC_HTTP_SERVER_PORT),

				// Read and write timeouts to avoid the server hang
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			},
		}
	}

	return server
}

// Start Starts listening and serving requests
func (s *Server) Start() {
	log.AddSimple(log.Info, "HTTP Server started, running on address "+s.Server.Addr)

	if err := s.Server.ListenAndServe(); err != nil {
		log.AddSimple(log.Error, "Error running API server listening on port "+s.Server.Addr)
	}
}

// registerHandlers Registers all the server handlers
func registerHandlers(router *mux.Router) {
	handlers.Healthcheck(router)
	handlers.Status(router)
	handlers.Log(router)
	handlers.Data(router)
	handlers.Updater(router)
}
