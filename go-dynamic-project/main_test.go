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
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	if contentType := res.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected application/json; got %v", contentType)
	}

	var fetchedTasks []Task
	if err := json.NewDecoder(res.Body).Decode(&fetchedTasks); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(fetchedTasks) < 2 {
		t.Errorf("expected at least 2 tasks; got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := Task{Title: "Test Task"}
	body, _ := json.Marshal(newTask)

	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created; got %v", res.StatusCode)
	}

	var createdTask Task
	if err := json.NewDecoder(res.Body).Decode(&createdTask); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if createdTask.Title != newTask.Title {
		t.Errorf("expected title %q; got %q", newTask.Title, createdTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		getTasksHandler(w, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
