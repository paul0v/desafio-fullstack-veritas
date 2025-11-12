package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// saveTasks writes current tasks map to the given filename as JSON array
func saveTasks(filename string) error {
	mu.Lock()
	arr := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		arr = append(arr, t)
	}
	mu.Unlock()

	b, err := json.MarshalIndent(arr, "", "  ")
	if err != nil {
		return err
	}
	// write to temp file then rename for atomicity
	tmp := filename + ".tmp"
	if err := ioutil.WriteFile(tmp, b, 0644); err != nil {
		return err
	}
	if err := os.Rename(tmp, filename); err != nil {
		return err
	}
	return nil
}

// loadTasks loads tasks from filename (if exists) and populates the in-memory map
func loadTasks(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil // nothing to load
	}
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	var arr []Task
	if err := json.Unmarshal(b, &arr); err != nil {
		return err
	}
	mu.Lock()
	tasks = make(map[int]Task)
	maxID := 0
	for _, t := range arr {
		tasks[t.ID] = t
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	nextID = maxID + 1
	mu.Unlock()
	log.Printf("loaded %d tasks from %s", len(arr), filename)
	return nil
}
