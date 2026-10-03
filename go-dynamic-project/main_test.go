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

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	var fetchedTasks []Task
	if err := json.NewDecoder(rr.Body).Decode(&fetchedTasks); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if len(fetchedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(fetchedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := Task{Title: "New Test Task"}
	body, _ := json.Marshal(newTask)

	req, err := http.NewRequest("POST", "/api/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addTasksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdTask Task
	if err := json.NewDecoder(rr.Body).Decode(&createdTask); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if createdTask.Title != newTask.Title {
		t.Errorf("expected task title %s, got %s", newTask.Title, createdTask.Title)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		getTasksHandler(rr, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
