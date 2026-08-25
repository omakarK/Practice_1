package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	w := httptest.NewRecorder()
	getTasksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}
	if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected application/json content-type, got %v", contentType)
	}
}

func TestAddTasksHandler(t *testing.T) {
	body := []byte(`{"title":"New Task"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	addTasksHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status Created, got %v", w.Code)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
