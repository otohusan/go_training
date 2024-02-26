package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello", "sa")

	var x = []int{2, 4, 6}

	var slice = append(x, 4)

	var i = 0
	for i < len(slice) {
		println(slice[i])
		i++

	}
}
