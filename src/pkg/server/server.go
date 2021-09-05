package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"pkg/data"
)

type Server struct {
	Router	*mux.Router
	Data	data.Data
}

func NewServer(data data.Data) *Server {
	router := mux.NewRouter().StrictSlash(true)
	router = NewAuthRouter(router, data)
	router = NewUserRouter(router, data)
	s := Server{
		Router: router,
		Data: data,
	}
	return &s
}

func (rcvr *Server) Start() {
	if rcvr.Data.Config.HTTPS == "on" {
		fmt.Println("Listening on port 8443")
		http.ListenAndServeTLS(":8443", rcvr.Data.Config.Cert, rcvr.Data.Config.Key, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	} else {
		fmt.Println("Listening on port", rcvr.Data.Config.ServerListenPort)
		http.ListenAndServe(rcvr.Data.Config.ServerListenPort, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	}
}

func (rcvr *Server) Init() {}