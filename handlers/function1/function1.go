package function1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"example.com/task3gcp/models"
	"example.com/task3gcp/utils"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GetAllEmployees", getAllEmployees)
}

// GetAllEmployees returns a list of all employees.
// @Summary Get all employees
// @Description Get a list of all employees
// @Produce json
// @Success 200 {array} Employee
// @Router /employees [get]
func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	utils.InitLogger()
	utils.InfoLog("Request is being Processed for GetAllEmployees")
	log.Print("Request received: GetAllEmployees")

	client, err := utils.CreateFirestoreClient()
	if err != nil {
		log.Print("Failed to create Firestore client:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	employees := []models.Employee{}
	iter := client.Collection("employees").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err != nil {
			log.Print("Error fetching document:", err)
			break
		}
		var employee models.Employee
		if err := doc.DataTo(&employee); err != nil {
			log.Print("Error parsing employee data:", err)
			continue
		}
		employees = append(employees, employee)
	}
	log.Print("Sending response: GetAllEmployees")
	respondWithJSON(w, http.StatusOK, employees)
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
