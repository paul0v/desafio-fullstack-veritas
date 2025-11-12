package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// try to load persisted tasks (if any)
	if err := loadTasks("data.json"); err != nil {
		log.Printf("warning: failed to load tasks: %v", err)
	}

	// list & create
	mux.HandleFunc("/tasks", tasksHandler)
	// operations on single task
	mux.HandleFunc("/tasks/", taskHandler)

	addr := ":8080"
	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, logRequest(mux)); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// logRequest is a middleware that logs basic request info
func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
