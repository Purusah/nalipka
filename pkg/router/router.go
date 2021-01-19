package router

import (
	"encoding/json"
	"net/http"
)

const headerContentType = "Content-Type"
const mimeHTML = "text/html; charset=UTF-8"
const mimeJSON = "application/json; charset=UTF-8"

type OkResponse struct {
	Data interface{} `json:"data"`
}

// HTTPErrorMethodNotAllowed ...
func HTTPErrorMethodNotAllowed(w http.ResponseWriter) {
	w.Header().Set(headerContentType, mimeHTML)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed"))
}

// HTTPErrorBadRequest ...
func HTTPErrorBadRequest(w http.ResponseWriter) {
	w.Header().Set(headerContentType, mimeHTML)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
}

// HTTPErrorNotFound ...
func HTTPErrorNotFound(w http.ResponseWriter) {
	w.Header().Set(headerContentType, mimeHTML)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

// HTTPOk ...
func HTTPOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set(headerContentType, mimeJSON)
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(OkResponse{Data: data})
}
