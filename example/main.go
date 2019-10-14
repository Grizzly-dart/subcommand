package main

import (
	"context"
	"flag"
	"fmt"
	subcommands "github.com/Grizzly-dart/subcommand"
	"os"
)

type HelloCmd struct {
	subcommands.CommandMixin
}

func (*HelloCmd) Name() string { return "hello" }

func (*HelloCmd) Synopsis() string { return "Greets!"}

func (*HelloCmd) Usage() string { return "hello" }

func (cmd *HelloCmd) Description() string {	return cmd.Synopsis() }

// Execute executes the command and returns an ExitStatus.
func (*HelloCmd) Execute(commander *subcommands.Commander, ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	fmt.Println("Hello!")

	return subcommands.ExitSuccess
}

func main() {
	subcommands.Register(&HelloCmd{}, "")
	subcommands.Register(&subcommands.HelpCommand{}, "")

	flag.Parse()
	os.Exit(int(subcommands.Execute(context.Background())))
}
