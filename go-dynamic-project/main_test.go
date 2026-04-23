package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasks(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
