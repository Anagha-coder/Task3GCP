package controller

import (
	"context"
	"encoding/json"

	// "log"
	"net/http"
	"strconv"
	"task3gcp/models"
	"task3gcp/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

var validate = validator.New()

// GetAllEmployees returns a list of all employees.
// @Summary Get all employees
// @Description Get a list of all employees
// @Produce json
// @Success 200 {array} Employee
// @Router /employees [get]
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	client, err := utils.CreateFirestoreClient()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	employees := []models.Employee{}
	iter := client.Collection("employees").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var employee models.Employee
		if err := doc.DataTo(&employee); err != nil {
			continue
		}
		employees = append(employees, employee)
	}

	respondWithJSON(w, http.StatusOK, employees)
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
// @Router /employees/{id} [get]
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	client, err := utils.CreateFirestoreClient()
	if err != nil {
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
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve employee from Firestore")
		return
	}

	var employee models.Employee
	if err := doc.DataTo(&employee); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to parse employee data")
		return
	}

	respondWithJSON(w, http.StatusOK, employee)
}

// SearchEmployees returns employees based on the specified field and value.
// @Summary Search employees by field and value
// @Description Search employees based on the specified field and value
// @Produce json
// @Param field query string true "Field to search (e.g., FirstName, LastName, Email, Role)"
// @Param value query string true "Value to search for"
// @Success 200 {array} Employee
// @Failure 400 "Bad Request: Invalid field or value"
// @Failure 500 "Internal Server Error"
// @Router /employees/search [get]
// SearchEmployee searches employees in Firestore based on specific fields and values.
func SearchEmployee(w http.ResponseWriter, r *http.Request) {
	// Parse search parameters from query string
	queryParams := r.URL.Query()

	// Retrieve search field and value from query parameters
	searchField := queryParams.Get("field")
	searchValue := queryParams.Get("value")

	if searchField == "" || searchValue == "" {
		respondWithError(w, http.StatusBadRequest, "Search field and value are required")
		return
	}

	client, err := utils.CreateFirestoreClient()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	// Define a query to search employees based on the specified field and value
	query := client.Collection("employees").Where(searchField, "==", searchValue)

	// Run the query
	iter := query.Documents(context.Background())

	var searchResults []models.Employee

	// Iterate through the query results and populate the searchResults slice
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve search results from Firestore")
			return
		}

		var employee models.Employee
		if err := doc.DataTo(&employee); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to parse employee data")
			return
		}

		searchResults = append(searchResults, employee)
	}

	respondWithJSON(w, http.StatusOK, searchResults)
}

// CreateEmployeeHandler creates a new employee.
// @Summary Create a new employee
// @Description Create a new employee
// @Accept json
// @Produce json
// @Param employee body Employee true "Employee object to be created"
// @Success 201 {object} map[string]string "Employee created successfully"
// @Failure 400 "Invalid request payload"
// @Failure 500 "Internal Server Error"
// @Router /employees [post]
func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate input data
	if err := validate.Struct(employee); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create a Firestore client
	client, err := utils.CreateFirestoreClient()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	// Read existing employees from Firestore (assuming you have a collection named "employees")
	iter := client.Collection("employees").Documents(context.Background())
	var existingEmployees []models.Employee
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to read employee data from Firestore")
			return
		}
		var emp models.Employee
		if err := doc.DataTo(&emp); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to parse employee data from Firestore")
			return
		}
		existingEmployees = append(existingEmployees, emp)
	}

	// Generate a unique ID for the new employee
	newEmployeeID := generateUniqueEmployeeID(existingEmployees)

	// Set the new employee ID
	employee.ID = newEmployeeID

	// Add the new employee to Firestore
	_, _, err = client.Collection("employees").Add(context.Background(), employee)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create employee in Firestore")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Employee created successfully"})
}

func generateUniqueEmployeeID(existingEmployees []models.Employee) int {
	highestID := 0

	// Find the highest existing ID from the Firestore data
	for _, employee := range existingEmployees {
		if employee.ID > highestID {
			highestID = employee.ID
		}
	}

	// Increment the highest existing ID to generate a new unique ID
	return highestID + 1
}

// UpdateEmployeeHandler updates an existing employee by ID.
// @Summary Update an existing employee
// @Description Update an existing employee by ID
// @Accept json
// @Produce json
// @Param id path number true "Employee ID to be updated"
// @Param employee body Employee true "Updated employee object"
// @Success 200 {object} map[string]string "Employee updated successfully"
// @Failure 400 "Invalid employee ID"
// @Failure 400 "Invalid request payload"
// @Failure 404 "Employee not found"
// @Failure 500 "Internal Server Error"
// @Router /employees/{id} [put]
// UpdateEmployee updates the employee details in Firestore based on the provided ID.
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	var updatedEmployee models.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate input data
	if err := validate.Struct(updatedEmployee); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	client, err := utils.CreateFirestoreClient()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	// Define a query to retrieve the document with the specified "ID" field value
	query := client.Collection("employees").Where("ID", "==", id).Limit(1)

	// Run the query
	iter := query.Documents(context.Background())
	doc, err := iter.Next()

	if err != nil {
		if status.Code(err) == codes.NotFound {
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve employee from Firestore")
		return
	}

	// Exclude ID field from the updated data
	updatedEmployee.ID = id

	_, err = doc.Ref.Set(context.Background(), updatedEmployee)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update employee in Firestore")
		return
	}

	respondWithJSON(w, http.StatusOK, updatedEmployee)
}

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
// @Router /employees/{id} [delete]
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	client, err := utils.CreateFirestoreClient()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create Firestore client")
		return
	}
	defer client.Close()

	// Define a query to retrieve the document with the specified "ID" field value
	query := client.Collection("employees").Where("ID", "==", id).Limit(1)

	// Run the query
	iter := query.Documents(context.Background())
	doc, err := iter.Next()

	if err != nil {
		if status.Code(err) == codes.NotFound {
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve employee from Firestore")
		return
	}

	_, err = doc.Ref.Delete(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete employee from Firestore")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Employee deleted successfully"})
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
