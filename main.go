package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "child" {
		childMain()
	} else {
		hostMain()
	}
}
