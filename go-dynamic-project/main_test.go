package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	getTasksHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

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
	body := []byte(`{"title":"New Task"}`)
	req := httptest.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created, got %v", resp.Status)
	}

	var newTask Task
	if err := json.NewDecoder(resp.Body).Decode(&newTask); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if newTask.Title != "New Task" {
		t.Errorf("expected title 'New Task', got %s", newTask.Title)
	}
}

func TestAddTasksHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status MethodNotAllowed, got %v", resp.Status)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
