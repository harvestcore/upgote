package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/harvestcore/HarvestCCode/src/config"
	"github.com/harvestcore/HarvestCCode/src/log"
)

var lock = &sync.Mutex{}

// Server HTTP server of the software
type Server struct {
	server *http.Server
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
			server: &http.Server{
				Handler: router,
				Addr:    ":" + config.GetManager().GetVariable(config.HCC_HTTP_SERVER_PORT),

				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			},
		}
	}

	return server
}

// Start Starts listening and serving requests
func (s *Server) Start() {
	port := config.GetManager().GetVariable(config.HCC_HTTP_SERVER_PORT)
	log.AddSimple(log.Info, "HTTP Server started, running on port "+port)
	s.server.ListenAndServe()
}

// registerHandlers Registers all the router handlers
func registerHandlers(router *mux.Router) {

}
