package main

import "fmt"

func main() {
	str := "Hèllo!"

	strrun := []rune(str)

	fmt.Println(str)

	fmt.Println(len(strrun))
	fmt.Println(len(str))

}
