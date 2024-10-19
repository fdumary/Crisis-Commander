package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Define the Plan struct
type Plan struct {
	Name        string
	Description string
	PlanID      string
	Category    string
}

var plansDB = make(map[string]Plan)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/plans", handlePlans) // GET all or DELETE all
	http.HandleFunc("/plan/", handlePlan)  // Handle GET, POST, DELETE requests

	http.ListenAndServe(":8080", nil) // Start the server on port 8080
}

// Handlers

func handlePlans(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		getPlans(writer, request)
	} else if request.Method == http.MethodDelete {
		deletePlans(writer, request)
	} else {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePlan(writer http.ResponseWriter, request *http.Request) {
	planID := strings.TrimPrefix(request.URL.Path, "/plan/")

	if request.Method == http.MethodGet {
		getPlanByID(writer, planID)
	} else if request.Method == http.MethodDelete {
		deletePlan(writer, planID)
	} else if request.Method == http.MethodPost {
		addPlan(writer, request)
	} else {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Controllers

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello user, welcome to our todo list app")
}

func getPlans(writer http.ResponseWriter, request *http.Request) {
	// Get the length of the map
	numPlans := len(plansDB)

	if numPlans == 0 {
		fmt.Fprintln(writer, "No plans available")
		return
	}

	fmt.Fprintf(writer, "You have %d plans:\n", numPlans)
	for id, plan := range plansDB {
		fmt.Fprintf(
			writer,
			"Plan ID: %s, Plan Name: %s, Plan Category: %s, Plan Description: \n%s\n",
			id, plan.Name, plan.Category, plan.Description)
	}
}

func deletePlans(writer http.ResponseWriter, request *http.Request) {
	// Clear the plansDB map
	plansDB = make(map[string]Plan)

	// Confirm deletion
	fmt.Fprintln(writer, "All plans have been deleted.")
}

func getPlanByID(writer http.ResponseWriter, planID string) {
	// Look up the plan in the map
	plan, exists := plansDB[planID]
	if !exists {
		http.Error(writer, "Plan not found", http.StatusNotFound)
		return
	}

	// Return the plan details
	fmt.Fprintf(
		writer,
		"Plan ID: %s, Plan Name: %s, Plan Category: %s, Plan Description: \n%s",
		planID, plan.Name, plan.Category, plan.Description)
}

// New function to delete a plan by ID
func deletePlan(writer http.ResponseWriter, planID string) {
	// Check if the plan exists
	_, exists := plansDB[planID]
	if !exists {
		http.Error(writer, "Plan not found", http.StatusNotFound)
		return
	}

	// Delete the plan
	delete(plansDB, planID)

	// Confirm deletion
	fmt.Fprintf(writer, "Plan with ID: %s has been deleted.\n", planID)
}

func addPlan(writer http.ResponseWriter, request *http.Request) {
	// Parse form data
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Get fields from form data
	name := strings.TrimSpace(request.FormValue("name"))
	description := strings.TrimSpace(request.FormValue("description"))
	category := strings.TrimSpace(request.FormValue("category"))

	// Validate fields
	if name == "" || description == "" || category == "" {
		http.Error(writer, "All fields are required", http.StatusBadRequest)
		return
	}

	// Generate a unique ID
	planID := uuid.New().String()

	// Create a new Plan object
	newPlan := Plan{
		Name:        name,
		Description: description,
		PlanID:      planID,
		Category:    category,
	}

	// Add the plan to the map with the generated ID
	plansDB[planID] = newPlan

	// Confirm plan addition
	// Return the plan details
	fmt.Fprintf(
		writer,
		"Plan ID: %s, Plan Name: %s, Plan Category: %s, Plan Description: \n%s",
		planID, newPlan.Name, newPlan.Category, newPlan.Description)
}
