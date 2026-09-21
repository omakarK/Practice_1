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
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var fetchedTasks []Task
	if err := json.NewDecoder(rec.Body).Decode(&fetchedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(fetchedTasks) == 0 {
		t.Errorf("expected non-empty task list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	body := bytes.NewBufferString(`{"title":"Test task"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", body)
	rec := httptest.NewRecorder()
	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rec.Code)
	}

	var createdTask Task
	if err := json.NewDecoder(rec.Body).Decode(&createdTask); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if createdTask.Title != "Test task" || createdTask.ID == 0 {
		t.Errorf("unexpected created task: %+v", createdTask)
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
