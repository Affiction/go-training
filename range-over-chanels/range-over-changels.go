package main

func main() {
	c := make(chan string, 5)

	c <- "a"
	c <- "b"
	close(c)

	for v := range c {
		println(v)
	}
}