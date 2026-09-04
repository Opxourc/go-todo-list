# go-todo-list

A CLI Todo List written purely in Go for demonstration purposes!

## Running

```text
go run .
```

Tasks are stored in memory and are removed when the program exits.

## Commands

| Command | Description |
| --- | --- |
| `add <due-date> <objective>` | Add a task. Dates use `YYYY-MM-DD`. |
| `list` | List all tasks. |
| `complete <id>` | Mark a task as complete. |
| `delete <id>` | Delete a task. |
| `edit <id> <due-date> <objective>` | Change a task's due date or objective. |
| `help` | Show the command list. |
| `quit` | Exit the program. |

Example:

```text
add 2026-12-31 Finish the project
list
complete 1
quit
```
