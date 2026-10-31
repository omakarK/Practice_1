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

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	var returnedTasks []Task
	if err := json.NewDecoder(res.Body).Decode(&returnedTasks); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(returnedTasks) < 2 {
		t.Errorf("expected at least 2 tasks, got %d", len(returnedTasks))
	}
}

func TestAddTasksHandler(t *testing.T) {
	t.Run("Valid POST", func(t *testing.T) {
		body := bytes.NewBufferString(`{"title": "Test Task"}`)
		req := httptest.NewRequest(http.MethodPost, "/api/tasks", body)
		w := httptest.NewRecorder()

		addTasksHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected status 201 Created, got %d", res.StatusCode)
		}

		var created Task
		if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
			t.Fatalf("failed to decode created task: %v", err)
		}

		if created.Title != "Test Task" {
			t.Errorf("expected title 'Test Task', got '%s'", created.Title)
		}
		if created.ID == 0 {
			t.Errorf("expected non-zero ID")
		}
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
		w := httptest.NewRecorder()

		addTasksHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("expected status 405 Method Not Allowed, got %d", res.StatusCode)
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		body := bytes.NewBufferString(`invalid json`)
		req := httptest.NewRequest(http.MethodPost, "/api/tasks", body)
		w := httptest.NewRecorder()

		addTasksHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status 400 Bad Request, got %d", res.StatusCode)
		}
	})
}

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		getTasksHandler(w, req)
	}
}

func BenchmarkGetTasksHandlerParallel(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}
