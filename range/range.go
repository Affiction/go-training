package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	myMap := map[string]string{"test1": "1", "test2": "2", "test3": "3"}

	for i, v := range myMap {
		fmt.Println(i, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
