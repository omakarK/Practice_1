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

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	var fetchedTasks []Task
	if err := json.Unmarshal(rr.Body.Bytes(), &fetchedTasks); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(fetchedTasks) == 0 {
		t.Errorf("expected non-empty task list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := map[string]string{"title": "Test Task"}
	body, _ := json.Marshal(newTask)

	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	addTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdTask Task
	if err := json.Unmarshal(rr.Body.Bytes(), &createdTask); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if createdTask.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got %v", createdTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
