package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"pkg"
	"pkg/configuration"
)

type Server struct {
	Router			*mux.Router
	Config			configuration.Configuration
	DbClient		*mongo.Client
	AuthService		root.AuthService
	UserService		root.UserService
}

func NewServer(config configuration.Configuration, dbClient *mongo.Client, dbService root.DbService, auth root.AuthService, user root.UserService) *Server {
	router := mux.NewRouter().StrictSlash(true)
	router = NewAuthRouter(router, config, dbClient, auth)
	router = NewUserRouter(router, config, dbClient, user)
	s := Server{
		Router: router,
		Config: config,
		DbClient: dbClient,
		AuthService: auth,
		UserService: user,
	}
	return &s
}

func (rcvr *Server) Start() {
	if rcvr.Config.HTTPS == "on" {
		fmt.Println("Listening on port 8443")
		http.ListenAndServeTLS(":8443", rcvr.Config.Cert, rcvr.Config.Key, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	} else {
		fmt.Println("Listening on port", rcvr.Config.ServerListenPort)
		http.ListenAndServe(rcvr.Config.ServerListenPort, handlers.LoggingHandler(os.Stdout, rcvr.Router))
	}
}

func (rcvr *Server) Init() {}