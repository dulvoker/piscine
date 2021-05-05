package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	prog_name := os.Args[0]
	casted := []rune(prog_name)
	for i := 0; i < len(prog_name); i++ {
		z01.PrintRune(casted[i])
	}
}
