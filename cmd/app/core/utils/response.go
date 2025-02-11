package utils

import (
	"encoding/json"
	"net/http"
)

type RequestDetails struct {
	RequestId string `json:"request_id"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Agent     string `json:"agent"`
}

type Response struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
	Request RequestDetails `json:"request"`
}

func SendResponse(w http.ResponseWriter, r *http.Request, data interface{}, message string, statusCode int) {
	// Retrieve requestID from context
	requestID, _ := r.Context().Value("requestID").(string)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Prepare the response structure
	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
		Request: RequestDetails{
			RequestId: requestID,
			Path:      r.URL.String(),
			Method:    r.Method,
			Agent:     r.UserAgent(),
		},
	}

	// Encode response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		SendError(w, r, "Failed to process response", http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	// Retrieve requestID from context
	requestID, _ := r.Context().Value("requestID").(string)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Prepare the error response structure
	response := Response{
		Status:  "error",
		Message: message,
		Request: RequestDetails{
			RequestId: requestID,
			Path:      r.URL.String(),
			Method:    r.Method,
			Agent:     r.UserAgent(),
		},
	}

	// Encode error response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle failure to encode error response
		http.Error(w, "Failed to process error response", http.StatusInternalServerError)
	}
}
