package main

import (
	"fmt"
	"interpreter/3-Lexer/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s!\nThis is G, the programming language by ingyu1008.\n", user)
	repl.Start(os.Stdin, os.Stdout)
}
