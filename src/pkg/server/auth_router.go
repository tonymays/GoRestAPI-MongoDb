package server

import (
//	"encoding/json"
//	"errors"
	"github.com/gorilla/mux"
//	"go.mongodb.org/mongo-driver/mongo"
//	"io"
//	"io/ioutil"
//	"net/http"
	"pkg/data"
)

type authRouter struct {
	data	data.Data
}

func NewAuthRouter(router *mux.Router, data data.Data) *mux.Router {
	//authRouter :=  authRouter{data}
	return router
}