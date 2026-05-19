package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func BenchmarkGetTasks(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		getTasksHandler(w, req)
	}
}

func BenchmarkGetTasksParallel(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}

func BenchmarkAddTasks(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		body := strings.NewReader(`{"title":"Benchmark task"}`)
		req := httptest.NewRequest("POST", "/api/tasks", body)
		w := httptest.NewRecorder()
		addTasksHandler(w, req)
	}
}
