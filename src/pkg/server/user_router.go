package server

import (
	"encoding/json"
//	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"pkg"
	"pkg/configuration"
	"pkg/db"
)

// ---- metricRouter ----
type metricRouter struct {
	config			configuration.Configuration
	db				*db.TestDataCache
	metricService	root.MetricService
}

// ---- NewMetricRouter ----
func NewMetricRouter(config configuration.Configuration, router *mux.Router, db *db.TestDataCache, metricService root.MetricService) *mux.Router {
	// fill in the metricRouter structure ... note we could have mongo clients,
	// db service clients, all kinds of things can now be passed in here
	metricRouter := metricRouter{config, db, metricService}

	/*
		I always add options because I never know if the client that is making
		the call requires them to be in place like React and Vuejs via Axios
	*/

	// Setup OPTIONS Method for all endpoints
	router.HandleFunc("/metric/{key}", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/metric/{key}/sum", HandleOptionsRequest).Methods("OPTIONS")
	router.HandleFunc("/metric/{key}/active", HandleOptionsRequest).Methods("OPTIONS")

	// Setup endpoint receivers - all enpoints must check themselves into the middleware for a COVID Test.
	router.HandleFunc("/metric/{key}", VerifyToken(metricRouter.postMetric, config)).Methods("POST")
	router.HandleFunc("/metric/{key}", VerifyToken(metricRouter.getMetric, config)).Methods("GET")
	router.HandleFunc("/metric/{key}", VerifyToken(metricRouter.clearOutdatedMetrics, config)).Methods("DELETE")
	router.HandleFunc("/metric/{key}/sum", VerifyToken(metricRouter.sumMetric, config)).Methods("GET")
	router.HandleFunc("/metric/{key}/active", VerifyToken(metricRouter.showActiveMetrics, config)).Methods("GET")

	/*
		return our router now that it is setup as a dependency of our server
		to which the server is a dependency of the core app
		This function is housed here to keep the endpoint definitions
		associated with their receivers
	*/
	return router
}

// ---- metricRouter.postMetric ----
func (rcvr *metricRouter) postMetric(w http.ResponseWriter, r *http.Request) {
	// grab the body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	var m root.Metric
	if err := json.Unmarshal(body, &m); err != nil {
		panic(err)
	}

	// grab the key
	vars := mux.Vars(r)
	m.Key = vars["key"]

	// add the metric
	m, err = rcvr.metricService.AddMetric(m)
	if err == nil {
		// respond with success
		w = SetResponseHeaders(w)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	} else {
		throw(w, err)
	}
}

// ---- metricRouter.getMetric ----
func (rcvr *metricRouter) getMetric(w http.ResponseWriter, r *http.Request) {
	m, err := rcvr.metricService.GetMetrics()
	if err == nil {
		w = SetResponseHeaders(w)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	} else {
		throw(w, err)
	}
}

// ---- metricRouter.showActiveMetric ----
func (rcvr *metricRouter) showActiveMetrics(w http.ResponseWriter, r *http.Request) {
	m, err := rcvr.metricService.ShowActiveMetrics()
	if err == nil {
		w = SetResponseHeaders(w)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	} else {
		throw(w, err)
	}
}

// ---- metricRouter.sumMetric ----
func (rcvr *metricRouter) sumMetric(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := rcvr.metricService.SumMetrics(vars["key"])
	if err == nil {
		w = SetResponseHeaders(w)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	} else {
		throw(w, err)
	}
}

// ---- metricRouter.clearOutdatedMetrics ----
func (rcvr *metricRouter) clearOutdatedMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := rcvr.metricService.ClearOutdatedMetrics(vars["key"])
	if err == nil {
		w = SetResponseHeaders(w)
		w.WriteHeader(http.StatusOK)
	} else {
		throw(w, err)
	}
}