package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var plans = make(map[string]string) // Initialize a map to store plans

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/plans", handlePlans) // Use the same route for both GET and POST
	http.HandleFunc("/plan/", handlePlan)  // Handle GET, DELETE requests

	http.ListenAndServe(":8080", nil) // Start the server on port 8080
}

// Handlers

func handlePlans(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		getPlans(writer, request)
	} else if request.Method == http.MethodPost {
		addPlan(writer, request)
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
	numPlans := len(plans)

	if numPlans == 0 {
		fmt.Fprintln(writer, "No plans available")
		return
	}

	fmt.Fprintf(writer, "You have %d plans:\n", numPlans)
	for id, plan := range plans {
		fmt.Fprintf(writer, "%s. %s\n", id, plan) // Display plan ID and description
	}
}

func getPlanByID(writer http.ResponseWriter, planID string) {
	// Look up the plan in the map
	plan, exists := plans[planID]
	if !exists {
		http.Error(writer, "Plan not found", http.StatusNotFound)
		return
	}

	// Return the plan details
	fmt.Fprintf(writer, "Plan ID: %s, Plan: %s\n", planID, plan)
}

// New function to delete a plan by ID
func deletePlan(writer http.ResponseWriter, planID string) {
	// Check if the plan exists
	_, exists := plans[planID]
	if !exists {
		http.Error(writer, "Plan not found", http.StatusNotFound)
		return
	}

	// Delete the plan
	delete(plans, planID)

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

	// Get the plan from form data
	newPlan := strings.TrimSpace(request.FormValue("plan"))

	if newPlan == "" {
		http.Error(writer, "Plan cannot be empty", http.StatusBadRequest)
		return
	}

	// Generate a unique ID
	planID := uuid.New().String()

	// Add the plan to the map with the generated ID
	plans[planID] = newPlan

	// Confirm plan addition
	fmt.Fprintf(writer, "Plan Created: %s (ID: %s)\n", newPlan, planID)
}
