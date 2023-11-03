package function5

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/task3gcp/utils"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteEmployeeHandler deletes an existing employee by ID.
// @Summary Delete an existing employee
// @Description Delete an existing employee by ID
// @Accept json
// @Produce json
// @Param id path number true "Employee ID to be deleted"
// @Success 200 {object} map[string]string "Employee deleted successfully"
// @Failure 400 "Invalid employee ID"
// @Failure 404 "Employee not found"
// @Failure 500 "Internal Server Error"
// @Router /function-5/{id} [delete]
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	utils.InitLogger()
	utils.InfoLog("Request is being Processed for DeleteEmployeeHandler")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print("Invalid employee ID:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	client, err := utils.CreateFirestoreClient()
	if err != nil {
		log.Print("Failed to create Firestore client:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	utils.InfoLog("Request received: DeleteEmployeeHandler")

	// Define a query to retrieve the document with the specified "ID" field value
	query := client.Collection("employees").Where("ID", "==", id).Limit(1)

	// Run the query
	iter := query.Documents(context.Background())
	doc, err := iter.Next()

	if err != nil {
		if status.Code(err) == codes.NotFound {
			utils.InfoLog("Employee not found")
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}
		log.Print("Failed to retrieve employee from Firestore:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve employee from Firestore")
		return
	}

	_, err = doc.Ref.Delete(context.Background())
	if err != nil {
		log.Print("Failed to delete employee from Firestore:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete employee from Firestore")
		return
	}

	utils.InfoLog("Employee deleted successfully")
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Employee deleted successfully"})
	utils.InfoLog("Response Sent")
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
