package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func open_and_read(s string) bool {
	file, err := os.Open(s)
	if err != nil {
		for _, char := range "ERROR: " {
			z01.PrintRune(char)
		}
		io.WriteString(os.Stdout, err.Error())
		z01.PrintRune('\n')
		return false
	}
	sad, _ := file.Stat()

	arr := make([]byte, sad.Size())
	file.Read(arr)
	for _, char := range arr {
		z01.PrintRune(rune(char))
	}
	file.Close()
	return true
}

func main() {
	if len(os.Args) > 1 {
		for _, each := range os.Args[1:] {
			if open_and_read(each) == false {
				os.Exit(1)
			}
		}
		z01.PrintRune('\n')
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	}
}
