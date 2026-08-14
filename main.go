package main

import (
	"fmt"
	"os"
	"strconv"	
)


func main() {
	todos := Todos{}
	todos.Load("todos.json")

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
		if len(todos) == 0 {
			fmt.Println("No tasks found! Add one with: todo add \"Your task\"")
			return
		}
		for _, task := range todos {
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
		todos.Add(title)
		todos.Save("todos.json")
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
		todos.Complete(id)
		todos.Save("todos.json")
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
		todos.Delete(id)
		todos.Save("todos.json")
		fmt.Println("Deleted task:", id)

	default:
		fmt.Println("Unknown command:", command)
	}
}
