package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var tasks = make(map[string]string) // Initialize a map to store tasks

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/tasks", handleTasks) // Use the same route for both GET and POST
	http.HandleFunc("/task/", handleTask)  // Handle GET, DELETE requests

	http.ListenAndServe(":8080", nil) // Start the server on port 8080
}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello user, welcome to our todo list app")
}

func handleTasks(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		getTasks(writer, request)
	} else if request.Method == http.MethodPost {
		addTask(writer, request)
	} else {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTask(writer http.ResponseWriter, request *http.Request) {
	taskID := strings.TrimPrefix(request.URL.Path, "/task/")

	if request.Method == http.MethodGet {
		getTaskByID(writer, taskID)
	} else if request.Method == http.MethodDelete {
		deleteTask(writer, taskID)
	} else {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(writer http.ResponseWriter, request *http.Request) {
	// Get the length of the map
	numTasks := len(tasks)

	if numTasks == 0 {
		fmt.Fprintln(writer, "No tasks available")
		return
	}

	fmt.Fprintf(writer, "You have %d tasks:\n", numTasks)
	for id, task := range tasks {
		fmt.Fprintf(writer, "%s. %s\n", id, task) // Display task ID and description
	}
}

func getTaskByID(writer http.ResponseWriter, taskID string) {
	// Look up the task in the map
	task, exists := tasks[taskID]
	if !exists {
		http.Error(writer, "Task not found", http.StatusNotFound)
		return
	}

	// Return the task details
	fmt.Fprintf(writer, "Task ID: %s, Task: %s\n", taskID, task)
}

// New function to delete a task by ID
func deleteTask(writer http.ResponseWriter, taskID string) {
	// Check if the task exists
	_, exists := tasks[taskID]
	if !exists {
		http.Error(writer, "Task not found", http.StatusNotFound)
		return
	}

	// Delete the task
	delete(tasks, taskID)

	// Confirm deletion
	fmt.Fprintf(writer, "Task with ID: %s has been deleted.\n", taskID)
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	// Parse form data
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Get the task from form data
	newTask := strings.TrimSpace(request.FormValue("task"))

	if newTask == "" {
		http.Error(writer, "Task cannot be empty", http.StatusBadRequest)
		return
	}

	// Generate a unique ID
	taskID := uuid.New().String()

	// Add the task to the map with the generated ID
	tasks[taskID] = newTask

	// Confirm task addition
	fmt.Fprintf(writer, "Task Created: %s (ID: %s)\n", newTask, taskID)
}
