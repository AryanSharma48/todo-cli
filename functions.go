package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt,omitempty"` // pointer to make CompletedAt an optional field
}

type Todos []Task

func (t* Todos) Add(title string) {
	task := Task{
		ID:        len(*t) + 1,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	*t = append(*t, task)
}

func (t *Todos) Complete(id int) error {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Done = true
			now := time.Now()
			(*t)[i].CompletedAt = &now
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

func (t *Todos) Delete(id int) error {
	for i, task := range *t {
		if task.ID == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

func (t *Todos) Undo(id int) error {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Done = false
			(*t)[i].CompletedAt = nil

			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

func (t *Todos) Clear() error {

	var remaining Todos
	for _, task := range *t {
		if !task.Done {
			remaining = append(remaining, task)
		}
	}

	clearedCount := len(*t) - len(remaining)
	if clearedCount == 0 {
		fmt.Println("No completed tasks found!")
    } else {
        fmt.Printf("Cleared %d completed task(s)!\n", clearedCount)
    }

	*t = remaining
	return nil
}