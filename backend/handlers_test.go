package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// helper to reset in-memory store between tests
func resetStore() {
	mu.Lock()
	tasks = make(map[int]Task)
	nextID = 1
	mu.Unlock()
}

func cleanupDataFile() {
	_ = os.Remove("data.json")
}

func TestCreateAndListTasks(t *testing.T) {
	resetStore()
	defer cleanupDataFile()

	// create a task
	payload := map[string]string{"title": "Test Task", "description": "desc"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	tasksHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d, body: %s", rr.Code, rr.Body.String())
	}
	var created Task
	if err := json.Unmarshal(rr.Body.Bytes(), &created); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}
	if created.ID == 0 || created.Title != "Test Task" {
		t.Fatalf("unexpected created task: %+v", created)
	}

	// list tasks
	req2 := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr2 := httptest.NewRecorder()
	tasksHandler(rr2, req2)
	if rr2.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr2.Code)
	}
	var arr []Task
	if err := json.Unmarshal(rr2.Body.Bytes(), &arr); err != nil {
		t.Fatalf("invalid json list: %v", err)
	}
	if len(arr) != 1 || arr[0].Title != "Test Task" {
		t.Fatalf("unexpected list content: %+v", arr)
	}
}

func TestValidateTitleRequired(t *testing.T) {
	resetStore()
	defer cleanupDataFile()

	// create with empty title
	payload := map[string]string{"title": "   ", "description": "x"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	tasksHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for empty title, got %d", rr.Code)
	}
	// ensure no tasks created
	mu.Lock()
	n := len(tasks)
	mu.Unlock()
	if n != 0 {
		t.Fatalf("expected 0 tasks, got %d", n)
	}
}

func TestUpdateTask(t *testing.T) {
	resetStore()
	defer cleanupDataFile()

	// create a task first
	payload := map[string]string{"title": "Initial", "description": "d"}
	b, _ := json.Marshal(payload)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	tasksHandler(rr, req)
	if rr.Code != http.StatusCreated {
		t.Fatalf("create failed: %d", rr.Code)
	}
	var created Task
	_ = json.Unmarshal(rr.Body.Bytes(), &created)

	// update
	upd := map[string]string{"title": "Updated", "description": "nd", "status": "in_progress"}
	ub, _ := json.Marshal(upd)
	req2 := httptest.NewRequest(http.MethodPut, "/tasks/"+itoa(created.ID), bytes.NewReader(ub))
	req2.Header.Set("Content-Type", "application/json")
	rr2 := httptest.NewRecorder()
	taskHandler(rr2, req2)
	if rr2.Code != http.StatusOK {
		t.Fatalf("update failed: %d", rr2.Code)
	}
	var up Task
	_ = json.Unmarshal(rr2.Body.Bytes(), &up)
	if up.Title != "Updated" || up.Status != "in_progress" {
		t.Fatalf("unexpected updated task: %+v", up)
	}
}

// helper itoa
func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}

// include fmt and io usages references to satisfy linters if needed
var _ = io.EOF
