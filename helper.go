package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func (t *Todos) PrintTable(filter string) {
	if len(*t) == 0 {
		fmt.Println("No tasks found! Add one with: todo add \"Your task\"")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATUS\tTASK\tCREATED\tCOMPLETED")
	fmt.Fprintln(w, "--\t------\t----\t-------\t---------")

	displayedCount := 0

	for _, task := range *t {
		
		if filter == "done" && !task.Done {
			continue // skip pending tasks
		}
		if filter == "pending" && task.Done {
			continue // skip completed tasks
		}

		displayedCount++

		status := "[ ]"
		completedStr := "-"
		if task.Done {
			status = "[x]"
			if task.CompletedAt != nil {
				completedStr = task.CompletedAt.Format("02 Jan 15:04")
			}
		}
		createdStr := task.CreatedAt.Format("02 Jan 15:04")
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", task.ID, status, task.Title, createdStr, completedStr)
	}

	w.Flush()

	if displayedCount == 0 {
		if filter == "done" {
			fmt.Println("No completed tasks found.")
		} else if filter == "pending" {
			fmt.Println("All caught up! No pending tasks.")
		}
	}
}


func (t* Todos) getNextID() int {
	maxID := 0
	for i := range *t {
		if (*t)[i].ID > maxID {
			maxID = (*t)[i].ID
		}
	}
	return maxID + 1
}