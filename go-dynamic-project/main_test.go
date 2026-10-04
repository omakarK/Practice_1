package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	getTasksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json; got %v", contentType)
	}
}

func TestAddTasksHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/api/tasks", nil)
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status BadRequest for empty body; got %v", w.Code)
	}
}
