package main

import "fmt"

func main() {
	fmt.Println("Comparison")
	fmt.Println(1 == 1)
	fmt.Println("" == "")
	fmt.Println("Hello" == "Hello")
	fmt.Println("Hello" == "hello")
	fmt.Println(1 == 2)
	fmt.Println(1.2 == 1.2)
	fmt.Println("Compare arrays", [2]int{} == [2]int{})
	fmt.Println("Compare arrays", [2]int{} == [2]int{1})
	fmt.Println("Compare arrays", [2]int{1,2} == [2]int{1,2})
	fmt.Println("Compare arrays", [2]int{1,2} == [2]int{1,3})

	fmt.Println("Array")

	var a [5]string
	fmt.Println(a, len(a), cap(a))

	var b = [5]int{1, 2, 3}
	fmt.Println(b, len(b), cap(b))

	b[1] = 12
	fmt.Println("Mutated Array", b)

	fmt.Println("Nested Array")
	var c = [2][3]int{}
	fmt.Println(c, len(c), cap(c))

	fmt.Println("Fill Array")
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c, len(c), cap(c))
}
