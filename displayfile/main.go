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
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("File name missing")
		}
		arr := make([]byte, 15)
		file.Read(arr)
		fmt.Println(string(arr))
	}
}
