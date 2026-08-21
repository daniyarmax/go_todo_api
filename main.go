package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	IsDone      bool   `json:"is_done"`
}

var tasks []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task.ID = len(tasks) + 1
	task.CreatedAt = time.Now().Format(time.RFC3339)

	tasks = append(tasks, task)

	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"message": "Task created successfully",
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetTasks(w, r)
		case http.MethodPost:
			CreateTask(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
