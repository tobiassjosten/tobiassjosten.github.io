// Package scripts holds the site's utility commands. Each command is an
// exported function with the signature func([]string) error and is registered
// in Commands so the root binary can dispatch to it: `go run . <command>`.
//
// The package is importable (not package main) so the test suite under tests/
// can reuse its content helpers.
package scripts

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/joho/godotenv"
)

// Command is one invocable utility, exposed as `go run . <Name>`.
type Command struct {
	Name    string
	Args    string
	Summary string
	Run     func(args []string) error
}

// Commands lists every available command in display order.
var Commands = []Command{
	{"analyze-goodreads", "[csv]", "List Goodreads books not yet in content/books/", AnalyzeGoodreads},
	{"create-content", "<type> [slug]", "Scaffold a new content file via `hugo new`", CreateContent},
	{"download-highlights", "[book-id]", "Print Readwise highlights for a book", DownloadHighlights},
	{"illustrate-thinkers", "", "Generate missing thinker portraits via OpenAI", IllustrateThinkers},
	{"import-goodreads", "[csv]", "Import books, authors and covers from a Goodreads export", ImportGoodreads},
	{"reverse-thinker-relationship", "", "Migrate book→thinker links into thinker files", ReverseThinkerRelationship},
}

// Run dispatches to the command named by args[0], passing the rest along. With
// no arguments it prints the list of available commands.
func Run(args []string) error {
	_ = godotenv.Load(".env.local", ".env")

	if len(args) == 0 {
		printCommands(os.Stdout)
		return nil
	}

	name := args[0]
	for _, command := range Commands {
		if command.Name == name {
			return command.Run(args[1:])
		}
	}

	fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", name)
	printCommands(os.Stderr)

	return fmt.Errorf("unknown command: %s", name)
}

func printCommands(w io.Writer) {
	fmt.Fprintln(w, "Usage: go run . <command> [arguments]")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Commands:")

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	for _, command := range Commands {
		fmt.Fprintf(tw, "  %s %s\t%s\n", command.Name, command.Args, command.Summary)
	}
	tw.Flush()
}
