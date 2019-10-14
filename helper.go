package subcommands

import (
	"context"
	"flag"
	"fmt"
)

// A HelpCommand is a Command implementing a "help" command for
// a given Commander.
// HelpCommand returns a Command which implements "help" for the
// DefaultCommander. Use Register(HelpCommand(), <group>) for it to be
// recognized.
type HelpCommand struct {
}

func (h *HelpCommand) Name() string           { return "help" }
func (h *HelpCommand) Synopsis() string       { return "describe subcommands and their syntax" }
func (h *HelpCommand) SetFlags(*flag.FlagSet) {}
func (h *HelpCommand) Usage() string {
	return `help [<subcommand>]`
}

func (h *HelpCommand) Description() string {
	return `With an argument, prints detailed information on the use of
	the specified subcommand. With no argument, print a list of
	all commands and a brief description of each.
	`
}

func (h *HelpCommand) Footer() string {
	return ""
}

func (h *HelpCommand) Execute(commander *Commander, _ context.Context, f *flag.FlagSet, args ...interface{}) ExitStatus {
	switch f.NArg() {
	case 0:
		DefaultCommander.Explain(commander.Output)
		return ExitSuccess

	case 1:
		for _, group := range commander.commands {
			for _, cmd := range group.commands {
				if f.Arg(0) != cmd.Name() {
					continue
				}
				DefaultCommander.ExplainCommand(commander.Output, cmd)
				return ExitSuccess
			}
		}
		fmt.Fprintf(commander.Error, "Subcommand %s not understood\n", f.Arg(0))
	}

	f.Usage()
	return ExitUsageError
}
