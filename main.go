package main

import (
	"fmt"
	"os"
	"os/user"
	"minotavar/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, this is the Minotavar programming language.\n" + user.Username)
	repl.Start(os.Stdin, os.Stdout)
}