package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Steins programming language Ψ(El Psy Congroo)!\n", user.Username)
	fmt.Print("Please enter the Steins program file to execute (or type commands directly): ")
	repl.Start(os.Stdin, os.Stdout)
}
