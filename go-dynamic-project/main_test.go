package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	rec := httptest.NewRecorder()
	getTasksHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", rec.Code)
	}
	expectedContentType := "application/json"
	if got := rec.Header().Get("Content-Type"); got != expectedContentType {
		t.Errorf("expected Content-Type %q, got %q", expectedContentType, got)
	}
}

func TestAddTasksHandler(t *testing.T) {
	// Simple test to ensure addTasksHandler works
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", nil)
	rec := httptest.NewRecorder()
	addTasksHandler(rec, req)
	// Body is empty so expects BadRequest
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status BadRequest, got %d", rec.Code)
	}
}
