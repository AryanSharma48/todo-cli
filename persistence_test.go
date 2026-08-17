package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "todos_test.json")

	// 1. Test Load on a non-existent file (should return nil and not error)
	todos := Todos{}
	err := todos.Load(tmpFile)
	if err != nil {
		t.Fatalf("Load failed for non-existent file: %v", err)
	}
	if len(todos) != 0 {
		t.Fatalf("expected empty todos, got %d", len(todos))
	}

	todos.Add("Task 1")
	todos.Add("Task 2")

	err = todos.Save(tmpFile)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify the file was indeed written
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to exist, but it doesn't", tmpFile)
	}

	loadedTodos := Todos{}
	err = loadedTodos.Load(tmpFile)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if len(loadedTodos) != len(todos) {
		t.Fatalf("expected %d todos, got %d", len(todos), len(loadedTodos))
	}

	if loadedTodos[0].Title != "Task 1" || loadedTodos[1].Title != "Task 2" {
		t.Errorf("loaded titles do not match: got %q and %q", loadedTodos[0].Title, loadedTodos[1].Title)
	}
}