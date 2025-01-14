package main

import (
	"fmt"
	"os"
	"os/user"
	"interpreter-in-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, this is the Minotavar programming language.\n" + user.Username)
	repl.Start(os.Stdin, os.Stdout)
}