package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) *List[T] {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}

	return lst
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	myMap := map[string]string{
		"foo":  "bar",
		"foo1": "bar1",
		"foo2": "bar2",
	}
	mapKeys := MapKeys(myMap)
	fmt.Printf("%v, type - %T\n", mapKeys, mapKeys)

	fmt.Println("List implementation --------------------------------")
	var myList List[int]

	myList.Push(1)
	myList.Push(2)
	myList.Push(3)

	fmt.Println("Get all:", myList.GetAll())
}
