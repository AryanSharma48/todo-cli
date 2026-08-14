package main

import (
	"fmt"
	"encoding/json"
	"os"
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
func (t *Todos) Save(filename string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func (t *Todos) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, t)
}