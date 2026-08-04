package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}
