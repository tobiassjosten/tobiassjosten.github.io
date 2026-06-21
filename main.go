package main

import (
	"fmt"
	"os"

	"tobiassjosten.net/scripts"
)

func main() {
	if err := scripts.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
