package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := 2
	var tasksReceived []Task
	if err := json.NewDecoder(rr.Body).Decode(&tasksReceived); err != nil {
		t.Fatal(err)
	}

	if len(tasksReceived) != expected {
		t.Errorf("handler returned unexpected body: got %v tasks want %v",
			len(tasksReceived), expected)
	}
}

func BenchmarkGetTasks(b *testing.B) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		getTasksHandler(rr, req)
	}
}

func BenchmarkGetTasksParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		req, _ := http.NewRequest("GET", "/api/tasks", nil)
		for pb.Next() {
			rr := httptest.NewRecorder()
			getTasksHandler(rr, req)
		}
	})
}
