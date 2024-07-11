package main

import "github.com/01-edu/z01"

func maina() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
