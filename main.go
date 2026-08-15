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
		fmt.Println("  undo <id>       - Mark task back to pending")
		fmt.Println("  del <id>        - Delete a task")
		fmt.Println("  clear           - Clear all completed tasks")
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		filter := ""
		if len(os.Args) >= 3 {
			switch os.Args[2] {
			case "--done":
				filter = "done"
			case "--pending":
				filter = "pending"
			default:
				fmt.Println("Unknown filter option! Use: todo list [--done|--pending]")
				return
			}
		}
		todos.PrintTable(filter)


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

	case "undo":
		if len(os.Args) < 3 {
			fmt.Println("Error: please specify a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		todos.Undo(id)
		todos.Save("todos.json")

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

	case "clear":
		if len(todos) == 0 {
			fmt.Println("No tasks found! Add one with: todo add \"Your task\"")
			return
		}
		todos.Clear()
		todos.Save("todos.json")

	default:
		fmt.Println("Unknown command:", command)
	}
}
