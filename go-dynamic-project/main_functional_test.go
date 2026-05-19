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
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var tasks []Task
	if err := json.NewDecoder(rr.Body).Decode(&tasks); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if len(tasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %v", len(tasks))
	}
}
