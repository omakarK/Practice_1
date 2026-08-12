package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	rec := httptest.NewRecorder()

	getTasksHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}
}

func TestAddTasksHandler(t *testing.T) {
	// Reset the state first
	mu.Lock()
	tasks = []Task{{ID: 1, Title: "Learn Go"}, {ID: 2, Title: "Build a dynamic website"}}
	nextID = 3
	mu.Unlock()

	taskData := map[string]string{"title": "Test Task"}
	bodyBytes, _ := json.Marshal(taskData)
	req := httptest.NewRequest("POST", "/api/tasks", bytes.NewBuffer(bodyBytes))
	rec := httptest.NewRecorder()

	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}

	var created Task
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if created.ID != 3 {
		t.Errorf("expected task ID 3, got %d", created.ID)
	}

	if created.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got '%s'", created.Title)
	}
}
