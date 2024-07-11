package main

import (
	"os"

	"github.com/01-edu/z01"
)

func mainr() {
	word := os.Args[:1]
	for _, char := range word {
		for _, words := range char {
			z01.PrintRune(words)
		}
	}
	z01.PrintRune('\n')
}
