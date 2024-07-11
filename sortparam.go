package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Sorttable(table []string) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func mainf() {
	word := os.Args[1:]
	Sorttable(word)
	for _, char := range word {
		for _, words := range char {
			z01.PrintRune(words)
		}
		z01.PrintRune('\n')
	}
}
