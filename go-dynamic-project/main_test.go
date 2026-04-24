package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	getTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var tasks []Task
	if err := json.NewDecoder(rr.Body).Decode(&tasks); err != nil {
		t.Fatal(err)
	}

	if len(tasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(tasks))
	}
}

func TestAddTask(t *testing.T) {
	task := Task{Title: "Test Task"}
	body, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	addTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var returnedTask Task
	if err := json.NewDecoder(rr.Body).Decode(&returnedTask); err != nil {
		t.Fatal(err)
	}

	if returnedTask.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got '%s'", returnedTask.Title)
	}
}

func BenchmarkGetTasks(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
