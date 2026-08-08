package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetTasksHandler(b *testing.B) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := httptest.NewRecorder()
			getTasksHandler(w, req)
		}
	})
}

func TestGetTasksHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()
	getTasksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var fetched []Task
	if err := json.Unmarshal(w.Body.Bytes(), &fetched); err != nil {
		t.Fatalf("failed to decode body: %v", err)
	}

	if len(fetched) == 0 {
		t.Error("expected non-empty tasks list")
	}
}

func TestAddTasksHandler(t *testing.T) {
	payload := []byte(`{"title":"New Task"}`)
	req := httptest.NewRequest("POST", "/api/tasks", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	addTasksHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}

	var fetched Task
	if err := json.Unmarshal(w.Body.Bytes(), &fetched); err != nil {
		t.Fatalf("failed to decode body: %v", err)
	}

	if fetched.Title != "New Task" || fetched.ID == 0 {
		t.Errorf("expected Title 'New Task' and non-zero ID, got %+v", fetched)
	}
}
