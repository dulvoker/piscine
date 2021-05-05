package piscine

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	prog_name := os.Args[0]
	casted := []rune(prog_name)
	for i := 0; i < len(prog_name); i++ {
		z01.PrintRune(rune(casted[i]))
	}
}
