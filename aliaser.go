package subcommands

// An aliaser is a Command wrapping another Command but returning a
// different name as its alias.
type aliaser struct {
	alias string
	Command
}

func (a *aliaser) Name() string { return a.alias }

// Alias returns a Command alias which implements a "commands" subcommand.
func Alias(alias string, cmd Command) Command {
	return &aliaser{alias, cmd}
}

// dealias recursivly dealiases a command until a non-aliased command
// is reached.
func dealias(cmd Command) Command {
	if alias, ok := cmd.(*aliaser); ok {
		return dealias(alias.Command)
	}

	return cmd
}
