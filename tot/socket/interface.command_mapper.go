package socket

/*
CommandMapper defines a management interface for the stack of pending commands.
*/
type CommandMapper interface {
	// Delete removes a command from the stack.
	Delete(commandID int)

	// Get retrieves a command from the stack.
	Get(commandID int) (Commander, error)

	// Set sets a command in the stack.
	Set(command Commander)
}
