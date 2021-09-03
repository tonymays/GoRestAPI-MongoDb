package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ---- Establish App struct ----
var ta App

// ---- setup ----
func setup() {
	// initialize the service environment for testing
	ta = App{}
	ta.Init("test")
}

// --- executeRequest ----
func executeRequest(ta App, req *http.Request) *httptest.ResponseRecorder {
	// Route Endpoint Executer Helper Function
	rr := httptest.NewRecorder()
	ta.Server.Router.ServeHTTP(rr, req)
	return rr
}

// ---- checkResponseCode ----
func checkResponseCode(t *testing.T, expected, actual int) {
	// Route Endpoint Response Check helper function
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// ---- Test 1: Get 200Ok from GET /metric/{key} ----
func TestGetMetric(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/metric/active_visitors", nil)
	req.Header.Add("Content-Type", "application/json")
	testResponse := executeRequest(ta, req)
	checkResponseCode(t, http.StatusOK, testResponse.Code)
}

// ---- Test 2: Get 200Ok from GET /metric/{key}/sum ----
func TestSumMetric(t *testing.T) {
	setup()
	req, _ := http.NewRequest("GET", "/metric/active_visitors/sum", nil)
	req.Header.Add("Content-Type", "application/json")
	testResponse := executeRequest(ta, req)
	checkResponseCode(t, http.StatusOK, testResponse.Code)
}

// ---- Test 3: Get 200Ok from POST /metric/{key} ----
func TestPostMetric(t *testing.T) {
	setup()
	payload := []byte(`{"value":25}`)
	req, _ := http.NewRequest("POST", "/metric/active_visitors", bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	testResponse := executeRequest(ta, req)
	checkResponseCode(t, http.StatusCreated, testResponse.Code)
}
