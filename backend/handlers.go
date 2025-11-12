package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	tasks  = make(map[int]Task)
	nextID = 1
	mu     sync.Mutex
)

// writeJSON writes a JSON response with given status code
func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// setCORS sets permissive CORS headers
func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// tasksHandler handles GET /tasks and POST /tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		arr := make([]Task, 0, len(tasks))
		for _, t := range tasks {
			arr = append(arr, t)
		}
		mu.Unlock()
		writeJSON(w, http.StatusOK, arr)
		return

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		var t Task
		if err := json.Unmarshal(body, &t); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		t.Title = strings.TrimSpace(t.Title)
		if t.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}
		if t.Status == "" {
			t.Status = StatusTodo
		} else if !ValidStatus(t.Status) {
			http.Error(w, "invalid status", http.StatusBadRequest)
			return
		}
		mu.Lock()
		t.ID = nextID
		nextID++
		tasks[t.ID] = t
		mu.Unlock()
		// persist
		if err := saveTasks("data.json"); err != nil {
			log.Printf("warning: failed to save tasks: %v", err)
		}
		writeJSON(w, http.StatusCreated, t)
		return

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// taskHandler handles GET/PUT/DELETE for /tasks/{id}
func taskHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// expect path like /tasks/{id}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		t, ok := tasks[id]
		mu.Unlock()
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, t)
		return

	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		var upd Task
		if err := json.Unmarshal(body, &upd); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		upd.Title = strings.TrimSpace(upd.Title)
		if upd.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}
		if upd.Status == "" {
			upd.Status = StatusTodo
		} else if !ValidStatus(upd.Status) {
			http.Error(w, "invalid status", http.StatusBadRequest)
			return
		}
		mu.Lock()
		_, ok := tasks[id]
		if !ok {
			mu.Unlock()
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		upd.ID = id
		tasks[id] = upd
		mu.Unlock()
		if err := saveTasks("data.json"); err != nil {
			log.Printf("warning: failed to save tasks: %v", err)
		}
		writeJSON(w, http.StatusOK, upd)
		return

	case http.MethodDelete:
		mu.Lock()
		_, ok := tasks[id]
		if ok {
			delete(tasks, id)
		}
		mu.Unlock()
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if err := saveTasks("data.json"); err != nil {
			log.Printf("warning: failed to save tasks: %v", err)
		}
		w.WriteHeader(http.StatusNoContent)
		return

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
