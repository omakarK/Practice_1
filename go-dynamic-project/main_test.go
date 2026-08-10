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
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var returnedTasks []Task
	if err := json.Unmarshal(rr.Body.Bytes(), &returnedTasks); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if len(returnedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %v", len(returnedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := Task{Title: "Test new task"}
	body, _ := json.Marshal(newTask)

	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var returnedTask Task
	if err := json.Unmarshal(rr.Body.Bytes(), &returnedTask); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if returnedTask.Title != "Test new task" {
		t.Errorf("expected task title to be 'Test new task', got %s", returnedTask.Title)
	}
}

func BenchmarkGetTasksParallel(b *testing.B) {
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
