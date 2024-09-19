package main

import (
	"net/http"
)

// ErrorHandler is a helper function for sending HTTP error responses
func ErrorHandler(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
}
