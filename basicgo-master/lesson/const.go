package main

import "fmt"

const Pi = "3.14"

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")
	const truth bool = true
	fmt.Println("Go rules?", truth)
}
