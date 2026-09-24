package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	var fetchedTasks []Task
	if err := json.Unmarshal(rr.Body.Bytes(), &fetchedTasks); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(fetchedTasks) == 0 {
		t.Errorf("expected non-empty task list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	taskPayload := []byte(`{"title": "Test Task"}`)
	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(taskPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	addTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var newTask Task
	if err := json.Unmarshal(rr.Body.Bytes(), &newTask); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if newTask.Title != "Test Task" {
		t.Errorf("expected task title 'Test Task', got '%s'", newTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
