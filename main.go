package main

import (
	"fmt"
	"log"
	"net/http"

	"task3gcp/controller"
	_ "task3gcp/docs" // Import generated docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title We're deploying REST API on GCP
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	// Serve Swagger documentation and UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/employees", controller.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", controller.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees/search/{field}/{value}", controller.SearchEmployee).Methods("GET")
	r.HandleFunc("/employees", controller.CreateEmployeeHandler).Methods("POST")
	r.HandleFunc("/employees/{id}", controller.UpdateEmployeeHandler).Methods("PUT")
	r.HandleFunc("/employees/{id}", controller.DeleteEmployeeHandler).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("Head over to http://localhost:8080/swagger/index.html to view Swagger documentation.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
