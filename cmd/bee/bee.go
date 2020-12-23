package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("useage: bee create|check [...]")
	}

	c, err := createCmd(os.Args[1:])
	if err != nil {
		panic(err)
	}

	if err = c.Run(); err != nil {
		panic(err)
	}
}
