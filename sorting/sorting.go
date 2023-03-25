package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"c", "b", "d", "e", "f", "g", "h", "a"}
	sort.Strings(str)
	fmt.Println("Sorting strings:", str)

	ints := []int{9, 9, 2134, 13, 2, 1}
	sort.Ints(ints)
	fmt.Println("Sorting ints:", ints)

	fmt.Println("Is strings sorted", sort.StringsAreSorted(str))
	fmt.Println("Is ints sorted", sort.IntsAreSorted(ints))
}
