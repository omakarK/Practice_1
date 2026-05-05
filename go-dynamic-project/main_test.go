package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Create a new recorder for each request to avoid race conditions
			// and simulate actual concurrent requests.
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}

func TestGetTasksHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	rr := httptest.NewRecorder()
	getTasksHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"Learn Go"},{"id":2,"title":"Build a dynamic website"}]`
	if rr.Body.String() != expected+"\n" && rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
