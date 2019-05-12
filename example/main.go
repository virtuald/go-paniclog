package main

import (
	"fmt"
	"os"

	"github.com/virtuald/go-paniclog"
)

func main() {

	f, err := os.Create("test.log")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	paniclog.RedirectStderr(f)
	f.Close()

	panic("this should end up in the file instead of the console")
}
