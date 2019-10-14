package subcommands

import (
	"context"
	"flag"
	"os"
	"path"
)

// DefaultCommander is the default commander using flag.CommandLine for flags
// and os.Args[0] for the command name.
var DefaultCommander *Commander

func init() {
	DefaultCommander = NewCommander(flag.CommandLine, path.Base(os.Args[0]))
}

// Register adds a subcommand to the supported subcommands in the
// specified group. (Help output is sorted and arranged by group
// name.)  The empty string is an acceptable group name; such
// subcommands are explained first before named groups. It is a
// wrapper around DefaultCommander.Register.
func Register(cmd Command, group string) {
	DefaultCommander.Register(cmd, group)
}

// ImportantFlag marks a top-level flag as important, which means it
// will be printed out as part of the output of an ordinary "help"
// subcommand.  (All flags, important or not, are printed by the
// "flags" subcommand.) It is a wrapper around
// DefaultCommander.ImportantFlag.
func ImportantFlag(name string) {
	DefaultCommander.ImportantFlag(name)
}

// Execute should be called once the default flags have been
// initialized by flag.Parse. It finds the correct subcommand and
// executes it, and returns an ExitStatus with the result. On a usage
// error, an appropriate message is printed to os.Stderr, and
// ExitUsageError is returned. The additional args are provided as-is
// to the Execute method of the selected Command. It is a wrapper
// around DefaultCommander.Execute.
func Execute(ctx context.Context, args ...interface{}) ExitStatus {
	return DefaultCommander.Execute(ctx, args...)
}
