package main

import "fmt"

type Base struct {
	num int
}

func (b Base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type Container struct {
	Base
	Str string
}

func main() {

	co := Container{
		Base: Base{
			num: 1,
		},
		Str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.Str)

	fmt.Println("also num:", co.Base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	fmt.Println("Mutation")
	fmt.Println("co.num", co.num)
	fmt.Println("co.Base.num", co.Base.num)

	co.Base.num = 12
	fmt.Println("co.num", co.num)
	fmt.Println("co.Base.num", co.Base.num)

	co.num = 42
	co.Base.num = 84
	fmt.Println("co.num", co.num)
	fmt.Println("co.Base.num", co.Base.num)
}
