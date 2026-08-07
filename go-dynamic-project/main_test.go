package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	getTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	var returnedTasks []Task
	if err := json.Unmarshal(rr.Body.Bytes(), &returnedTasks); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if len(returnedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(returnedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	taskPayload, _ := json.Marshal(Task{Title: "Write tests"})
	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(taskPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	addTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var returnedTask Task
	if err := json.Unmarshal(rr.Body.Bytes(), &returnedTask); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if returnedTask.Title != "Write tests" {
		t.Errorf("expected task title to be 'Write tests', got %s", returnedTask.Title)
	}
}

func TestConcurrentHandlers(t *testing.T) {
	var wg sync.WaitGroup
	workers := 20

	// Concurrent reads and writes
	for i := 0; i < workers; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", "/api/tasks", nil)
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}()

		go func() {
			defer wg.Done()
			taskPayload, _ := json.Marshal(Task{Title: "Concurrent task"})
			req, _ := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(taskPayload))
			rr := httptest.NewRecorder()
			addTasksHandler(rr, req)
		}()
	}

	wg.Wait()
}

func BenchmarkGetTasksHandler(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		req, _ := http.NewRequest("GET", "/api/tasks", nil)
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
