package app

import (
	"net/http"
	"encoding/json"
	"log"
	"gyg/middleware/config"
	"github.com/kr/pretty"
	"io"
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

	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", "attachment; filename=\"results.json\"")
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD");
	w.WriteHeader(http.StatusOK)

        searchTerm := r.URL.Query().Get("query")
	s.dispatcher.dispatch(searchTerm)

	pretty.Println("Request:" + r.URL.String())

	// Locking on Aggregated data
	data := <-output

	pretty.Println(data)

	err := json.NewEncoder(w).Encode(data)
	io.Copy(w, jso.en)
	if err != nil {
		pretty.Println("Server:" + err.Error())
	}
}