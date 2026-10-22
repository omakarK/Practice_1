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
		t.Fatalf("failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	getTasksHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var fetchedTasks []Task
	if err := json.Unmarshal(rec.Body.Bytes(), &fetchedTasks); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	if len(fetchedTasks) < 2 {
		t.Errorf("expected at least 2 initial tasks, got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	payload := []byte(`{"title": "Test task"}`)
	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	var newTask Task
	if err := json.Unmarshal(rec.Body.Bytes(), &newTask); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if newTask.Title != "Test task" {
		t.Errorf("expected task title 'Test task', got '%s'", newTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}
