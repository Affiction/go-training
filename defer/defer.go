package main

import (
	"fmt"
	"os"
)

func createFile(fileName string) *os.File {
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(file *os.File) {
	fmt.Fprintln(file, "data")
}

func closeFile(f *os.File) {
	err := f.Close()

	if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
	}
}

func main() {
	file := createFile("defer.txt")
	writeFile(file)
	defer closeFile(file)
}
