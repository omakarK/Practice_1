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

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK (200), got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", contentType)
	}

	var returnedTasks []Task
	if err := json.NewDecoder(res.Body).Decode(&returnedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(returnedTasks) < 2 {
		t.Fatalf("expected at least 2 tasks, got %d", len(returnedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := map[string]string{"title": "Test Task"}
	body, _ := json.Marshal(newTask)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewReader(body))
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected status Created (201), got %d", res.StatusCode)
	}

	var created Task
	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if created.Title != "Test Task" {
		t.Fatalf("expected title 'Test Task', got %q", created.Title)
	}
	if created.ID == 0 {
		t.Fatalf("expected non-zero task ID")
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
