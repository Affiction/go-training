package main

import (
	"fmt"
	"os"
)

func main() {
	panic("some error happened")

	fmt.Println("After panic")

	_, err := os.Create("/panic/file")
	if err != nil {
			panic(err)
	}
}