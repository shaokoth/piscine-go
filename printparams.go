package main

import (
	"os"

	"github.com/01-edu/z01"
)

func mainpo() {
	word := os.Args[1:]
	for _, str := range word {
		for _, strr := range str {
			z01.PrintRune(strr)
		}
	}
	z01.PrintRune('\n')
}
