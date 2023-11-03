package function2

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/task3gcp/models"
	"example.com/task3gcp/utils"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	functions.HTTP("GetEmployeeByID", getEmployeeByID)
}

// GetEmployeeByID returns a specific employee by ID.
// @Summary Get an employee by ID
// @Description Get an employee by ID
// @Produce json
// @Param id path number true "ID"
// @Success 200 {object} Employee
// @Failure 400 "Invalid employee ID"
// @Failure 404 "Employee not found"
// @Failure 500 "Internal Server Error"
// @Router /function-2/{id} [get]
func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	utils.InitLogger()
	log.Print("Request is being Processed for GetEmployeeByID")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print("Invalid employee ID:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	log.Print("Request received: GetEmployeeByID, ID:", id)

	client, err := utils.CreateFirestoreClient()
	if err != nil {
		log.Print("Failed to create Firestore client:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	// Define a query to retrieve the document with the specified "Id" field value
	query := client.Collection("employees").Where("ID", "==", id)

	// Run the query
	iter := query.Documents(context.Background())
	doc, err := iter.Next()

	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Print("Employee not found:", err)
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}
		log.Print("Failed to retrieve employee from Firestore:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve employee from Firestore")
		return
	}

	var employee models.Employee
	if err := doc.DataTo(&employee); err != nil {
		log.Print("Failed to parse employee data:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to parse employee data")
		return
	}

	log.Print("Sending response: GetEmployeeByID")
	respondWithJSON(w, http.StatusOK, employee)
	log.Print("Response Sent: GetEmployeeByID")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
