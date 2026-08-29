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
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
