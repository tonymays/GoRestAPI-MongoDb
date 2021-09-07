package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ---- jsonErr ----
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// ---- HandleOptionsRequest ----
func HandleOptionsRequest(w http.ResponseWriter, r *http.Request) {
	// establish what the options endpoints can handle
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD,GET,DELETE,POST,PATCH,PUT")
	w.WriteHeader(http.StatusOK)
}

// ---- SetResponseHeaders ----
func SetResponseHeaders( w http.ResponseWriter, authToken string, apiKey string ) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Auth-Token, API-Key")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD,GET,DELETE,POST,PATCH,PUT")
	if authToken != "" {
		w.Header().Add("Auth-Token", authToken)
	}
	if apiKey != "" {
		w.Header().Add("API-Key", apiKey)
	}
	return w
}

// ---- GetIpAddress ----
func GetIpAddress(r *http.Request) string {
	parts := strings.Split(r.RemoteAddr, ":")
	return parts[0]
}

// ---- throw ----
func throw(w http.ResponseWriter, callErr error) {
	w = SetResponseHeaders(w, "", "")
	w.WriteHeader(http.StatusForbidden)
	err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusForbidden, Text: callErr.Error()})
	if err != nil {
		panic(err)
	}
}