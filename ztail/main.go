package main

import (
	"fmt"
	"os"
)

func TrimAtoi(s string) int {
	casted := []rune(s)
	firstnumbcheck := 0
	var table []rune
	signcheck := 0
	for i := 0; i < len(s); i++ {
		if casted[i] > 47 && casted[i] < 58 {
			if firstnumbcheck == 1 {
				table = append(table, rune(casted[i]))
			} else if casted[i] != 47 {
				table = append(table, rune(casted[i]))
				firstnumbcheck = 1
			}
		}
		if casted[i] == 45 && len(table) == 0 {
			signcheck = 1
		}
	}
	sum := 0
	tens := 1
	for i := 0; i < len(table); i++ {
		for z := 0; z < len(table)-i-1; z++ {
			tens = tens * 10
		}
		sum = sum + (int(table[i])-48)*tens
		tens = 1
	}
	if signcheck == 1 {
		return -sum
	}
	return sum
}

func open_and_read(s string, how_much int) string {
	file, err := os.Open(s)
	if err != nil {
		fmt.Printf("open %v: no such file or directory", s)
		fmt.Printf("\n")
		return ""
	}
	sad, _ := file.Stat()
	arr := make([]byte, sad.Size())
	if sad.Size() > int64(how_much) {
		file.Read(arr)
		new_array := make([]byte, how_much)
		new_array = arr[sad.Size()-int64(how_much):]
		ret_array := string(new_array)
		file.Close()
		return ret_array
	}
	file.Read(arr)
	ret_array := string(arr)
	file.Close()
	return ret_array
}

func main() {
	if len(os.Args) > 3 {
		errors := 0
		needed := string(os.Args[1])
		if needed[0] == '-' && needed[1] == 'c' {
			number := TrimAtoi(os.Args[2])
			to_introduce := false
			if len(os.Args) > 4 {
				to_introduce = true
			}
			for i, each := range os.Args[3:] {
				if open_and_read(each, number) == "" {
					errors++
				} else {
					if i != 0 {
						fmt.Printf("\n")
					}
					if to_introduce == true {
						fmt.Printf("==> %v <==\n", each)
					}
					fmt.Printf("%v", open_and_read(each, number))
				}
			}
			if errors > 0 {
				os.Exit(1)
			}
		}
	}
}
