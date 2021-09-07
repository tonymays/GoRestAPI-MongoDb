package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ---- App Struct ----
var ta App

// ---- setup ----
func setup() {
	// put together a test environment
	ta = App{}
	ta.Init("test")
}

// ---- executeRequest ----
func executeRequest(ta App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ta.Server.Router.ServeHTTP(rr, req)
	return rr
}

// ---- checkResponseCode ----
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}