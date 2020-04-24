package main

import (
	"fmt"
	"os"

	"github.com/jdxj/study-flag/flag"
)

var usableCmd = `usable cmd:
	greet
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usableCmd)
		return
	}

	greet := &flag.GreetCmd{}
	flag.RegisterCmd(greet)

	subCmd := os.Args[1]
	err := flag.Handle(subCmd)
	if err == flag.ErrCmdHandlerNotFound {
		fmt.Print(usableCmd)
	} else if err != nil {
		panic(err)
	}
}
