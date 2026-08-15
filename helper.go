package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func (t *Todos) PrintTable() {
	if len(*t) == 0 {
		fmt.Println("No tasks found! Add one with: todo add \"Your task\"")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tSTATUS\tTASK\tCREATED\tCOMPLETED")
	fmt.Fprintln(w, "--\t------\t----\t-------\t---------")

	for _, task := range *t {
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