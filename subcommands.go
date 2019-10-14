package subcommands

import (
	"flag"
	"fmt"
	"io"
	"sort"
	"strings"
)

// A commandGroup represents a set of commands about a common topic.
type CommandGroup struct {
	name     string
	commands []Command
}

// Name returns the group name
func (g *CommandGroup) Name() string {
	return g.name
}

// Sorting of a slice of command groups.
type byGroupName []*CommandGroup

// TODO Sort by function rather than implement sortable?
func (p byGroupName) Len() int           { return len(p) }
func (p byGroupName) Less(i, j int) bool { return p[i].name < p[j].name }
func (p byGroupName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// explain prints a brief description of all the subcommands and the
// important top-level flags.
func (cdr *Commander) explain(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s <flags> <subcommand> <subcommand args>\n\n", cdr.name)
	sort.Sort(byGroupName(cdr.commands))
	for _, group := range cdr.commands {
		cdr.ExplainGroup(w, group)
	}
	if cdr.topFlags == nil {
		fmt.Fprintln(w, "\nNo top level flags.")
		return
	}

	sort.Strings(cdr.important)
	if len(cdr.important) == 0 {
		fmt.Fprintf(w, "\nUse \"%s flags\" for a list of top-level flags\n", cdr.name)
		return
	}

	fmt.Fprintf(w, "\nTop-level flags (use \"%s flags\" for a full list):\n", cdr.name)
	for _, name := range cdr.important {
		f := cdr.topFlags.Lookup(name)
		if f == nil {
			panic(fmt.Sprintf("Important flag (%s) is not defined", name))
		}
		fmt.Fprintf(w, "  -%s=%s: %s\n", f.Name, f.DefValue, f.Usage)
	}
}

// Sorting of the commands within a group.
func (g CommandGroup) Len() int           { return len(g.commands) }
func (g CommandGroup) Less(i, j int) bool { return g.commands[i].Name() < g.commands[j].Name() }
func (g CommandGroup) Swap(i, j int)      { g.commands[i], g.commands[j] = g.commands[j], g.commands[i] }

// explainGroup explains all the subcommands for a particular group.
func explainGroup(w io.Writer, group *CommandGroup) {
	if len(group.commands) == 0 {
		return
	}
	if group.name == "" {
		fmt.Fprintf(w, "Subcommands:\n")
	} else {
		fmt.Fprintf(w, "Subcommands for %s:\n", group.name)
	}
	sort.Sort(group)

	aliases := make(map[string][]string)
	for _, cmd := range group.commands {
		if alias, ok := cmd.(*aliaser); ok {
			root := dealias(alias).Name()

			if _, ok := aliases[root]; !ok {
				aliases[root] = []string{}
			}
			aliases[root] = append(aliases[root], alias.Name())
		}
	}

	for _, cmd := range group.commands {
		if _, ok := cmd.(*aliaser); ok {
			continue
		}

		name := cmd.Name()
		names := []string{name}

		if a, ok := aliases[name]; ok {
			names = append(names, a...)
		}

		fmt.Fprintf(w, "\t%-15s  %s\n", strings.Join(names, ", "), cmd.Synopsis())
	}
	fmt.Fprintln(w)
}

// explainCmd prints a brief description of a single command.
func explain(w io.Writer, cmd Command) {
	fmt.Fprintf(w, "Usage: %s\r\n\r\n", cmd.Usage())
	fmt.Fprintf(w, "%s\r\n", cmd.Description())
	subflags := flag.NewFlagSet(cmd.Name(), flag.PanicOnError)
	subflags.SetOutput(w)
	cmd.SetFlags(subflags)
	subflags.PrintDefaults()
	fmt.Fprintf(w, "%s\r\n", cmd.Footer())
}
