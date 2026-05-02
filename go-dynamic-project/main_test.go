package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api/tasks", nil)
		w := httptest.NewRecorder()
		getTasksHandler(w, req)
	}
}

func BenchmarkGetTasksParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest("GET", "/api/tasks", nil)
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
