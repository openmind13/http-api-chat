package apiserver

import (
	"encoding/json"
	"net/http"
)

// send response with error
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respondJSON(w, r, code, map[string]string{"error": err.Error()})
}

// sernd response
func (s *server) respondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
