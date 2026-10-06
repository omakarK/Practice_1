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
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}
	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := Task{Title: "Test Task"}
	body, _ := json.Marshal(newTask)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

	var createdTask Task
	if err := json.NewDecoder(rec.Body).Decode(&createdTask); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if createdTask.Title != "Test Task" {
		t.Errorf("expected task title 'Test Task', got '%s'", createdTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		getTasksHandler(rec, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}
