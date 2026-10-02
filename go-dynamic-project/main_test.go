package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()
	getTasksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200 OK; got %v", w.Code)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		req := httptest.NewRequest("GET", "/api/tasks", nil)
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
