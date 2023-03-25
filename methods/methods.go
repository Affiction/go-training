package main

import "fmt"

type Rect struct {
    Width, Height int
}

func (r *Rect) area() int {
    return r.Width * r.Height
}

func (r Rect) perim() int {
    return 2*r.Width + 2*r.Height
}

type CustomType struct {
	int
}

func main() {
    r := Rect{Width: 10, Height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perimeter:", r.perim())

		// By reference
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perimeter:", rp.perim())

		ct := CustomType{12}
		fmt.Printf("%#v\n",ct)
		fmt.Println(ct.int)
}