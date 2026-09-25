package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var (
	tasks  = []Task{{ID: 1, Title: "Learn Go"}, {ID: 2, Title: "Build a dynamic website"}}
	nextID = 3
	mu     sync.RWMutex
)

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Bolt Optimization: Use RLock to make a shallow copy of tasks, then encode JSON outside the lock.
	// This allows concurrent parallel reads and minimizes critical section duration.
	mu.RLock()
	tasksCopy := make([]Task, len(tasks))
	copy(tasksCopy, tasks)
	mu.RUnlock()

	json.NewEncoder(w).Encode(tasksCopy)
}

func addTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTasksHandler(w, r)
		} else if r.Method == http.MethodPost {
			addTasksHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
