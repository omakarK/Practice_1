package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	rec := httptest.NewRecorder()

	getTasksHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var fetched []Task
	if err := json.NewDecoder(rec.Body).Decode(&fetched); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(fetched) == 0 {
		t.Errorf("expected non-empty tasks list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	payload := []byte(`{"title":"New Task"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()

	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status Created, got %d", rec.Code)
	}

	var created Task
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode created task: %v", err)
	}
	if created.Title != "New Task" {
		t.Errorf("expected title 'New Task', got '%s'", created.Title)
	}
}

func TestAddTasksHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	rec := httptest.NewRecorder()

	addTasksHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status MethodNotAllowed, got %d", rec.Code)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}
