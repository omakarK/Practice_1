package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	rec := httptest.NewRecorder()

	getTasksHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK, got %d", rec.Code)
	}
	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected application/json, got %s", contentType)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		rec := httptest.NewRecorder()
		for pb.Next() {
			getTasksHandler(rec, req)
		}
	})
}
