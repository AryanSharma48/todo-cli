package main

import (
	"fmt"
	"encoding/json"
	"os"
	"strconv"	
)

type Task struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Done   bool    `json:"done"`
}

var tasks = []Task{}

func AddTask(newTask Task){
	tasks = append(tasks, newTask)
}

func CompleteTask(id int){
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return
		}
	}
}

func DeleteTask(id int){
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
		if os.IsNotExist(err){
			return nil
		}

		return err
	}


	return json.Unmarshal(data, &tasks)
}

func main() {
	LoadTasks("todos.json")

	if len(os.Args) < 2 {
		fmt.Println("Commands:")
		fmt.Println("  list            - Show all tasks")
		fmt.Println("  add <title>     - Add a new task")
		fmt.Println("  done <id>       - Mark task as completed")
		fmt.Println("  del <id>        - Delete a task")
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks found! Add one with: todo add \"Your task\"")
			return
		}
		for _, task := range tasks {
			status := "[ ]"
			if task.Done {
				status = "[x]"
			}
			fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
		}

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: please specify a task title.")
			return
		}
		title := os.Args[2]
		newTask := Task{
			ID:    len(tasks) + 1,
			Title: title,
			Done:  false,
		}
		AddTask(newTask)
		SaveTasks("todos.json")
		fmt.Println("Added task:", title)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: please specify a task ID.")
			return
		}
		// Convert string ID ("2") to int (2)
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		CompleteTask(id)
		SaveTasks("todos.json")
		fmt.Println("Marked task as done:", id)

	case "del":
		if len(os.Args) < 3 {
			fmt.Println("Error: please specify a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		DeleteTask(id)
		SaveTasks("todos.json")
		fmt.Println("Deleted task:", id)

	default:
		fmt.Println("Unknown command:", command)
	}
}
