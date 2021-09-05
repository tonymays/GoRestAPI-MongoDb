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

type userRouter struct {
	data	data.Data
}

func NewUserRouter(router *mux.Router, data data.Data) *mux.Router {
	//userRouter :=  userRouter{data}
	return router
}