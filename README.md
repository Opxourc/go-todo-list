# go-todo-list

`go-todo-list` is an interactive command-line todo manager written in Go. It
supports creating, viewing, editing, completing, and deleting tasks. Tasks are
loaded from disk when the application starts and saved back to disk when it
exits.

## Requirements

- Go 1.27 or newer

The application has no third-party dependencies. Run all commands from the
repository root.

## Run the application

Start the interactive prompt with:

```sh
go run .
```

The application creates `data/tasks.json` automatically if it does not exist.
The data file is relative to the current working directory, so start the
program from the repository root to use the project's normal data location.

The program displays a welcome message and then waits for one command per
line. Commands are case-insensitive, and blank lines are ignored. Type `help`
at any time to display the built-in command summary.

## Features

- Add tasks with a due date and free-form objective.
- Assign each task a unique positive integer ID during the current data set's
 lifetime.
- List tasks with their ID, due date, status, and objective.
- Mark tasks as complete.
- Edit a task's due date and objective without losing its completion status.
- Delete tasks by ID.
- Persist tasks as readable, indented JSON in `data/tasks.json`.
- Create the `data` directory and data file automatically when needed.
- Validate dates, IDs, argument counts, and unknown commands with clear errors.

## Commands

| Command | Syntax | Description |
| --- | --- | --- |
| Add | `add <due-date> <objective>` | Create a pending task. |
| List | `list` | Display all tasks, or `No tasks.` when none exist. |
| Complete | `complete <id>` | Mark the task as complete. |
| Delete | `delete <id>` | Remove the task permanently. |
| Edit | `edit <id> <due-date> <objective>` | Replace a task's due date and objective. |
| Help | `help` | Display the available commands. |
| Quit | `quit` | Exit the application and save tasks. |

### Arguments

- `<due-date>` must use the `YYYY-MM-DD` format, for example `2026-12-31`.
- `<id>` must be a positive integer matching an existing task.
- `<objective>` may contain multiple words. Everything after the due date is
 joined into the task objective.
- `edit` requires all three arguments. It preserves the task's current
 complete/pending status.

## Example session

```text
--- Todo List ---
Manage tasks that need to be completed!
Type 'help' to get a list for a list of commands.

> add 2026-12-31 Finish the project
Added task 1.
> add 2026-10-15 Review pull requests
Added task 2.
> list
1 | 2026-12-31 | pending | Finish the project
2 | 2026-10-15 | pending | Review pull requests
> complete 1
Completed task 1.
> edit 2 2026-11-01 Review open pull requests
Updated task 2.
> delete 1
Deleted task 1.
> quit
Successfully saved all tasks!
```

The prompt itself is not printed by the program; the `>` characters above are
only used to show user input clearly.

## Persistence

Tasks are kept in memory while the program runs and are written to
`data/tasks.json` when the program exits normally with `quit` or when standard
input reaches EOF. The file is read at startup, so tasks from earlier sessions
remain available.

The stored JSON contains the task fields used by the application:

```json
[
 {
  "ID": 1,
  "DueDate": "2026-12-31T00:00:00Z",
  "Objective": "Finish the project",
  "Complete": false
 }
]
```

Do not edit the file while the application is running. Invalid JSON causes the
application to stop with an error when it starts.

## Run tests

Run the full test suite with:

```sh
go test ./...
```

The repository currently includes tests for creating the data directory and
file, loading saved JSON, ignoring whitespace-only data files, and truncating
old contents when saving.

## Project layout

```text
.
├── main.go                 Application entry point and input loop
├── commands/               Implementations of add, list, complete, delete, edit, and help
├── database/               JSON file loading, saving, and persistence tests
├── models/                 Task and command types
├── util/                   Date, ID, and task lookup helpers
├── data/tasks.json         Runtime task storage, created automatically
├── go.mod                  Go module definition
└── LICENSE                 Project license
```

## License

See [LICENSE](LICENSE) for the license that applies to this project.
