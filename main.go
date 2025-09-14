package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get current user")
		os.Exit(1)
	}
	fmt.Printf("ハローワールド🌏\n%s !\n", usr.Username)

	repl.Start(os.Stdin, os.Stdout)
}
