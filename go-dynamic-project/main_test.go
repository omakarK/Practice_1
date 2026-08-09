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
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if rr.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("handler returned unexpected content-type: got %v want %v", rr.Header().Get("Content-Type"), expectedContentType)
	}

	var responseTasks []Task
	if err := json.Unmarshal(rr.Body.Bytes(), &responseTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(responseTasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(responseTasks))
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req, _ := http.NewRequest("GET", "/api/tasks", nil)
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
