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

	http.ListenAndServe(":8080", nil)
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

func getTasks(writer http.ResponseWriter, request *http.Request) {
	// Get the length of the map
	numTasks := len(tasks)

	if numTasks == 0 {
		fmt.Fprintln(writer, "No tasks available")
		return
	}

	fmt.Fprintf(writer, "You have %d tasks:\n", numTasks)
	for id, task := range tasks {
		fmt.Fprintf(writer, "%s. %s\n", id, task)
	}
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

	//generate ID
	var taskID = uuid.New().String()

	// Add the task to the map with an ID
	tasks[taskID] = newTask

	// Confirm task addition
	fmt.Fprintf(writer, "Task Created: %s\n", newTask)
}
