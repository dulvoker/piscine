package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		file, err := os.Open(string(os.Args[1]))
		if err != nil {
			fmt.Println("File name missing")
		}
		var table []byte
		file.Read(table)
		fmt.Println(string(table))
	}
}
