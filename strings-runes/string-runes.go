package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "Hello"
	b := "สวัสดี"

	fmt.Println(a, len(a))
	fmt.Println(b, len(b))

	fmt.Println("'a' iteration")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c \n", a[i])
	}

	fmt.Println("'b' iteration (length slice of bytes)")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%c \n", b[i])
	}

	fmt.Println("'a' Rune count:", utf8.RuneCountInString(a))
	fmt.Println("'b' Rune count:", utf8.RuneCountInString(b))

	for i := 0; i < utf8.RuneCountInString(b); i++ {
		fmt.Printf("%c \n", b[i])
	}

	fmt.Println("Range iteration")
	for idx, runeValue := range a {
		fmt.Printf("%c starts at %d\n", runeValue, idx)
	}
}
