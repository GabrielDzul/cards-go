package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// t := triangle{base: 12.2, height: 10.5}
	// s := square{base: 10, height: 10.0}

	// printArea(t)
	// printArea(s)

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

}
