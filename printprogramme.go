package main

import (
	"os"

	"github.com/01-edu/z01"
)

func mainpp() {
	word := os.Args[0]
	for _, name := range word[2:] {
		z01.PrintRune(name)
	}
	z01.PrintRune('\n')
}
