package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func variadic(nums ...int) int {
	var total int = 0
	for _, n := range nums {
		total += n
	}
	return total
}

func func1(nums ...interface{}) {
	for _, n := range nums {
		fmt.Println("Printing from func1", n)
	}
}

func main() {
	fmt.Println(sum(1, 3))
	a := []int{3, 4, 5, 6}
	fmt.Println(variadic(a...))

	func1(1, 2, "432", "hi")
}
