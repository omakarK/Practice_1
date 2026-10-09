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
	getTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var fetchedTasks []Task
	if err := json.Unmarshal(rr.Body.Bytes(), &fetchedTasks); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	if len(fetchedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	taskData := map[string]string{"title": "Test Task"}
	body, _ := json.Marshal(taskData)

	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	addTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var addedTask Task
	if err := json.Unmarshal(rr.Body.Bytes(), &addedTask); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	if addedTask.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got %s", addedTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
