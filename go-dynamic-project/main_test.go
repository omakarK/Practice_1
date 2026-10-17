package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	getTasksHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", resp.Status)
	}

	if contentType := resp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var fetchedTasks []Task
	if err := json.NewDecoder(resp.Body).Decode(&fetchedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(fetchedTasks) == 0 {
		t.Errorf("expected non-empty task list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	body := `{"title":"Test task"}`
	req := httptest.NewRequest("POST", "/api/tasks", strings.NewReader(body))
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created, got %v", resp.Status)
	}

	var newTask Task
	if err := json.NewDecoder(resp.Body).Decode(&newTask); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if newTask.Title != "Test task" {
		t.Errorf("expected title 'Test task', got %s", newTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		getTasksHandler(w, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
