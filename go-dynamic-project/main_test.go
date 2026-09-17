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
	rec := httptest.NewRecorder()

	getTasksHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", rec.Code)
	}

	var fetchedTasks []Task
	if err := json.NewDecoder(rec.Body).Decode(&fetchedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(fetchedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := map[string]string{"title": "Test Task"}
	body, _ := json.Marshal(newTask)
	req := httptest.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status Created, got %d", rec.Code)
	}
}

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
