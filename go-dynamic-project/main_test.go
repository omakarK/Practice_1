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
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected application/json, got %s", contentType)
	}

	var fetchedTasks []Task
	if err := json.Unmarshal(rec.Body.Bytes(), &fetchedTasks); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(fetchedTasks) < 2 {
		t.Fatalf("expected at least 2 tasks, got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := Task{Title: "Test Task"}
	body, _ := json.Marshal(newTask)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	addTasksHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}

	var createdTask Task
	if err := json.Unmarshal(rec.Body.Bytes(), &createdTask); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if createdTask.ID == 0 || createdTask.Title != "Test Task" {
		t.Fatalf("unexpected task returned: %+v", createdTask)
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
