package main

import (
	"fmt"
	"log"
	"os"
)

const filePath = "reading-files/data.txt"

func checkError(e error) {
	if e != nil {
		log.Fatalln("Error:", e.Error())
	}
}

func main() {
	// Basic file reading
	data, err := os.ReadFile(filePath)
	checkError(err)
	fmt.Println(string(data))

	f, err := os.Open(filePath)
	checkError(err)

	b1 := make([]byte, 4)
	n1, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0)
	checkError(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
}
