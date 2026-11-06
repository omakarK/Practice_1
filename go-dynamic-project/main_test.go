package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	w := httptest.NewRecorder()

	getTasksHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var returnedTasks []Task
	if err := json.NewDecoder(res.Body).Decode(&returnedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(returnedTasks) < 2 {
		t.Errorf("expected at least 2 initial tasks, got %d", len(returnedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	newTask := map[string]string{"title": "Test Task"}
	body, _ := json.Marshal(newTask)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewReader(body))
	w := httptest.NewRecorder()

	addTasksHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201 Created, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var created Task
	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if created.Title != "Test Task" || created.ID <= 0 {
		t.Errorf("unexpected created task: %+v", created)
	}
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Body.Reset()
		getTasksHandler(w, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
		w := httptest.NewRecorder()
		for pb.Next() {
			w.Body.Reset()
			getTasksHandler(w, req)
		}
	})
}

func BenchmarkAddTasksHandler(b *testing.B) {
	payload := []byte(`{"title":"Benchmark Task"}`)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Body.Reset()
		req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewReader(payload))
		addTasksHandler(w, req)
	}
}
