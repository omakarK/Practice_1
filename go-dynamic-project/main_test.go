package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	// Let's create a dummy slice of tasks to avoid empty slice edge cases
	tasks = []Task{
		{ID: 1, Title: "Learn Go"},
		{ID: 2, Title: "Build a dynamic website"},
	}

	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rec := httptest.NewRecorder()
			getTasksHandler(rec, req)
		}
	})
}
