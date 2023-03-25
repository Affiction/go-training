package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int { return len(s) }
func (s ByLength) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}
func (s ByLength) Less(a, b int) bool {
	return len(s[a]) < len(s[b])
}

func main() {
	slc := []string{"a", "aa", "aaa", "aaaa", "a"}
	sort.Sort(ByLength(slc))
	
	fmt.Println(slc)
}
