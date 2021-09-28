package main

import (
	"Beans/beans/pkg/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Beans programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
