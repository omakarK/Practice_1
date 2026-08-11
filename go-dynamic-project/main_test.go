package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content-type: got %v want %v", contentType, "application/json")
	}

	var returnedTasks []Task
	err = json.NewDecoder(rr.Body).Decode(&returnedTasks)
	if err != nil {
		t.Fatalf("could not parse returned JSON: %v", err)
	}

	if len(returnedTasks) != len(tasks) {
		t.Errorf("returned wrong number of tasks: got %v want %v", len(returnedTasks), len(tasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	body := `{"title":"New Task"}`
	req, err := http.NewRequest("POST", "/api/tasks", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTasksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content-type: got %v want %v", contentType, "application/json")
	}

	var returnedTask Task
	err = json.NewDecoder(rr.Body).Decode(&returnedTask)
	if err != nil {
		t.Fatalf("could not parse returned JSON: %v", err)
	}

	if returnedTask.Title != "New Task" {
		t.Errorf("returned task has wrong title: got %v want %v", returnedTask.Title, "New Task")
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
