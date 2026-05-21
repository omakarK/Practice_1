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
	if err := json.Unmarshal(rr.Body.Bytes(), &tasks); err != nil {
		t.Errorf("handler returned unexpected body: %v", err)
	}

	if len(tasks) == 0 {
		t.Errorf("handler returned empty task list")
	}
}
