package app

import (
	"net/http"
	"encoding/json"
	"log"
	"gyg/middleware/config"
	"github.com/kr/pretty"
)

type (
	Server struct {
		dispatcher *Dispatcher
	}
)

func NewServer(dispatcher *Dispatcher) *Server {
	return &Server{
		dispatcher: dispatcher,
	}
}

func (s *Server) Start() {
	listenAddr := config.Configuration.Server.Host + config.Configuration.Server.Port

	http.HandleFunc(config.Configuration.Server.Path, handler(s))

	// Should be last line as it is a blocking.
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func handler(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(r, w)
	}
}

func (s *Server) respond(r *http.Request, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

        searchTerm := r.URL.Query().Get("query")
	s.dispatcher.dispatch(searchTerm)
	// Locking on Aggregated data
	data := <-output

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		pretty.Println("Server:" + err.Error())
	}
}