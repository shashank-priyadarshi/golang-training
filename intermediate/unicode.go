package main

import (
	"fmt"
	"unicode"
)

type myString string

func main() {
	s := myString("abcXY@5±c")
	fmt.Println(s)
	fmt.Println(changeCase(s))
}

func changeCase(s myString) myString {
	for i, c := range s {
		if unicode.IsLetter(c) {
			if unicode.IsUpper(c) {
				c = unicode.ToLower(c)
			} else {
				c = unicode.ToUpper(c)
			}
			s = s[:i] + myString(string(c)) + s[i+1:]
		}
	}
	return s
}
