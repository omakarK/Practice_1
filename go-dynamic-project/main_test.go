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
	if err := json.Unmarshal(rr.Body.Bytes(), &fetchedTasks); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}
	if len(fetchedTasks) == 0 {
		t.Errorf("expected non-empty task list")
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
	if err := json.Unmarshal(rr.Body.Bytes(), &createdTask); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}
	if createdTask.ID == 0 || createdTask.Title != "New Test Task" {
		t.Errorf("unexpected created task: %+v", createdTask)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
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
