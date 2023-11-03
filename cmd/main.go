package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	_ "task3gcp/docs" // Import generated docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Employee Management System UI
// @version 1.0
// @description Google Cloud Platform to serve Cloud functions seamlessly
// @host us-central1-task3gcp.cloudfunctions.net
func main() {
	r := mux.NewRouter()

	// Serve Swagger documentation and UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/function-1", forwardRequest("https://us-central1-task3gcp.cloudfunctions.net/function-1", http.MethodGet)).Methods(http.MethodGet)
	r.HandleFunc("/function-2/{id}", forwardRequest("https://us-central1-task3gcp.cloudfunctions.net/function-2", http.MethodGet)).Methods(http.MethodGet)
	r.HandleFunc("/function-3", forwardRequest("https://us-central1-task3gcp.cloudfunctions.net/function-3", http.MethodPost)).Methods(http.MethodPost)
	r.HandleFunc("/function-4/{id}", forwardRequest("https://us-central1-task3gcp.cloudfunctions.net/function-4", http.MethodPut)).Methods(http.MethodPut)
	r.HandleFunc("/function-5/{id}", forwardRequest("https://us-central1-task3gcp.cloudfunctions.net/function-5", http.MethodDelete)).Methods(http.MethodDelete)

	http.Handle("/", r)

	fmt.Println("Head over to http://localhost:8085/swagger/index.html to view Swagger documentation.")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func forwardRequest(targetURL, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new request to the Cloud Function URL
		req, err := http.NewRequest(method, targetURL, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Copy headers from the original request to the new request
		for key, value := range r.Header {
			req.Header[key] = value
		}

		// Send the request to the Cloud Function URL
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copy the response from the Cloud Function to the original response writer
		for key, value := range resp.Header {
			w.Header()[key] = value
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}

}
