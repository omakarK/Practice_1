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
	// mu uses sync.RWMutex to allow concurrent reads while maintaining exclusive writes.
	// This improves performance for read-heavy workloads like getTasksHandler.
	mu sync.RWMutex
)

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	// ⚡ BOLT: Using RLock to allow multiple concurrent readers.
	// This reduces lock contention when many users are fetching tasks simultaneously.
	mu.RLock()
	defer mu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
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
