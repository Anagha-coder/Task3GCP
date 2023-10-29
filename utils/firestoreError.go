package utils

import (
	"net/http"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleFirestoreError(err error) (int, bool) {
	if status.Code(err) == codes.NotFound {
		// Firestore document not found error
		return http.StatusNotFound, false
	}

	if err == iterator.Done {
		// Iterator has reached the end of the results
		return http.StatusNotFound, false
	}

	// For other Firestore errors, you might want to log the error for debugging purposes
	// log.Printf("Firestore error: %v", err)

	// Internal server error for other Firestore errors
	return http.StatusInternalServerError, true
}
