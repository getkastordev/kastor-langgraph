package main

import (
	"fmt"
	"os"

	"github.com/getkastordev/kastor-langgraph/internal/plugin"
	protocol "github.com/weirdGuy/kastor/protocol/v1"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "serve" {
		fmt.Fprintln(os.Stderr, "usage: kastor-langgraph serve")
		os.Exit(2)
	}
	if err := protocol.Serve(os.Stdin, os.Stdout, plugin.Handler{}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
