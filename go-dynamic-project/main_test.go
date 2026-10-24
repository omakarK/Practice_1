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
	w := httptest.NewRecorder()

	getTasksHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", resp.Status)
	}

	if contentType := resp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %v", contentType)
	}

	var fetchedTasks []Task
	if err := json.NewDecoder(resp.Body).Decode(&fetchedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(fetchedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	body := []byte(`{"title": "Test Task"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created, got %v", resp.Status)
	}

	var createdTask Task
	if err := json.NewDecoder(resp.Body).Decode(&createdTask); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if createdTask.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got %s", createdTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
