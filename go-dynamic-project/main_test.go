package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var tasksResponse []Task
	if err := json.NewDecoder(rr.Body).Decode(&tasksResponse); err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	if len(tasksResponse) < 2 {
		t.Errorf("expected at least 2 tasks, got %v", len(tasksResponse))
	}
}

func TestAddTasksHandler(t *testing.T) {
	initialCount := len(tasks)

	taskJSON := []byte(`{"title": "New Task"}`)
	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	if len(tasks) != initialCount+1 {
		t.Errorf("expected task count to be %v, got %v", initialCount+1, len(tasks))
	}
}

func BenchmarkGetTasks(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
