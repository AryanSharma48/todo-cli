package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	todos := Todos{}

	// Add a task
	taskTitle := "Test Task 1"
	todos.Add(taskTitle)

	//Check: Was the task added?
	if len(todos) != 1 {
		t.Fatalf("expected 1 task, got %d", len(todos))
	}

	//Check: Are the properties correct?
	if todos[0].Title != taskTitle {
		t.Errorf("expected title %q, got %q", taskTitle, todos[0].Title)
	}

	if todos[0].Done != false {
		t.Errorf("expected task to be pending (done=false), got %v", todos[0].Done)
	}

	if todos[0].ID != 1 {
		t.Errorf("expected task ID to be 1, got %d", todos[0].ID)
	}
}

func TestComplete(t *testing.T) {
	todos := Todos{}
	todos.Add("Task to complete")

	// Success case: complete existing task
	err := todos.Complete(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !todos[0].Done {
		t.Errorf("expected task to be done, but it was not")
	}
	if todos[0].CompletedAt == nil {
		t.Errorf("expected CompletedAt to be set, but got nil")
	}

	// Error case: try to complete a non-existent task
	err = todos.Complete(999)
	if err == nil {
		t.Errorf("expected error for non-existent ID 999, but got nil")
	}
}

func TestDelete(t *testing.T) {
	todos := Todos{}
	todos.Add("Task to delete")

	err := todos.Delete(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(todos) != 0 {
		t.Fatalf("expected no task, got %d", len(todos))
	}

	// Error case: try to delete a non-existent task
	err = todos.Delete(999)
	if err == nil {
		t.Errorf("expected error for non-existent ID 999, but got nil")
	}

}