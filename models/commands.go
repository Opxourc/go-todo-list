package models

// Commands that a user can use.
type Command string

const (
	CommandAdd      Command = "add"
	CommandList     Command = "list"
	CommandComplete Command = "complete"
	CommandDelete   Command = "delete"
	CommandEdit     Command = "edit"
	CommandHelp     Command = "help"
	CommandQuit     Command = "quit"
)
