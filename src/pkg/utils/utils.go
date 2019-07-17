package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// JSONContentType is the MIME Type for JSON.
	JSONContentType = "application/json"
)

// Utility functions.

// WriteStatus200 writes the HTTP 200 status code.
func WriteStatus200(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// WriteStatus204 writes the HTTP 204 status code.
func WriteStatus204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// SetContentType sets the Content-Type header to contentType.
func SetContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

// SetContentJSON sets the content type to JSON.
func SetContentJSON(w http.ResponseWriter) {
	SetContentType(w, JSONContentType)
}

// WritePayload writes the bytestream to the response writer with 200 Status Code.
func WritePayload(w http.ResponseWriter, payload []byte) {
	WriteStatus200(w)
	w.Write(payload)
}

// WriteSerializeJSON serializes an interace and writes
func WriteSerializeJSON(w http.ResponseWriter, rawPayload interface{}) {
	serializedPayload, err := json.Marshal(rawPayload)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(serializedPayload))
	SetContentJSON(w)
	WritePayload(w, serializedPayload)

}
