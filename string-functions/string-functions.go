package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("my string", "my"))
	fmt.Println(strings.Count("my string", "y"))
	fmt.Println(strings.HasPrefix("my string", "my"))
	fmt.Println(strings.Index("aaabc", "b"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("foo", "o", "0", -1))
	fmt.Println(strings.Replace("foo", "o", "0", 1))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
}
