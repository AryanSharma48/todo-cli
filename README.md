# todo-cli

A simple, fast, and robust command-line todo list application written in Go. It stores your tasks locally in a JSON file, providing persistence across terminal sessions.

## Installation

Ensure you have Go installed, then clone the repository and build the binary:

```bash
git clone https://github.com/aryansharma48/todo-cli.git>
cd to-do
go build -o todo
```

You can then move the `todo` binary to a location in your PATH (e.g., `/usr/local/bin/todo`) to use it anywhere.

## Usage

Run the tool without arguments to see available commands:

```bash
todo
```

### Commands

* `todo list` - Show all tasks (both pending and completed).
* `todo list --pending` - Show only pending tasks.
* `todo list --done` - Show only completed tasks.
* `todo add <title>` - Add a new task to your list.
* `todo done <id>` - Mark a task as completed.
* `todo undo <id>` - Mark a completed task back as pending.
* `todo del <id>` - Delete a task entirely.
* `todo clear` - Remove all completed tasks from the list.

## Development

The project includes a Makefile for easy testing and linting.

* Run tests: `make test`
* Run linter: `make lint` (Requires golangci-lint)
