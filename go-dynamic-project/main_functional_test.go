package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	getTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var tasksResponse []Task
	if err := json.NewDecoder(rr.Body).Decode(&tasksResponse); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	// Initial tasks are defined in main.go
	if len(tasksResponse) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(tasksResponse))
	}
}
