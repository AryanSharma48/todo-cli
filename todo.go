package main

import (
	"encoding/json"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{}

func AddTask(newTask Task) {
	tasks = append(tasks, newTask)
}

func CompleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return
		}
	}
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

func SaveTasks(filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadTasks(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &tasks)
}
