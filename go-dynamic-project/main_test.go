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
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var gotTasks []Task
	if err := json.NewDecoder(rec.Body).Decode(&gotTasks); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(gotTasks) == 0 {
		t.Errorf("expected non-empty tasks list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	body := []byte(`{"title":"New Task"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	var createdTask Task
	if err := json.NewDecoder(rec.Body).Decode(&createdTask); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if createdTask.Title != "New Task" {
		t.Errorf("expected task title 'New Task', got %s", createdTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}
