package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int {
	return rec.width * rec.length
}

func main() {
	rec := rectangle{3, 4}
	// a := area(rec)
	a := rec.area()
	fmt.Println(a)
}
