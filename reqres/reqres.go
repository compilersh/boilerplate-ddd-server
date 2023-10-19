// reqres package provides a simple way to make HTTP requests and get responses.
package reqres

import (
	"encoding/json"
	"net/http"
)

// Validator is an interface that can be implemented by types that need to be validated.
type Validator interface {
	Validate() error
}

func DecodeJSON(r *http.Request, v Validator) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ResJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
