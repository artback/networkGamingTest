package webservice

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithError responds with a friendly JSON error
func RespondWithError(w http.ResponseWriter, r *http.Request, code int, message string) (int, error) {
	log.Print(r.Host, " ", message)
	return RespondWithJSON(w, code, message)
}

// RespondWithJSON responds with some JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) (int, error) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return w.Write(response)
}
