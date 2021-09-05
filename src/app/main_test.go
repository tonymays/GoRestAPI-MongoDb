package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ta App

func setup() {
	ta = App{}
	ta.Init("test")
}

func executeRequest(ta App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ta.Server.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}