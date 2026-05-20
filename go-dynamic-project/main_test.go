package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		getTasksHandler(rr, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
