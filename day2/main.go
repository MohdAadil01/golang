package main

import "fmt"

func main() {
	var students [4]string
	fmt.Println(len(students))

	array := [5]int{2, 4, 5, 7, 9}

	var i int = 0
	var j int = len(array) - 1

	var target int = 4

	// !applying binary search to find element
	for i <= j {
		var mid int = (i + j) / 2
		if array[mid] == target {
			fmt.Println("found at index", mid)
			break
		} else if array[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
}
