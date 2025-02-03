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

	// !SLICES
	var a = make([]int, 1, 4)
	fmt.Println(a)
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	var b = []int{2, 3, 4, 5, 56, 6}
	fmt.Println(b)

	for n := range 4 {
		b = append(b, n)
	}
	fmt.Println(b)

	mpp := make(map[string]int)

	mp := map[string]int{"bcd": 2}

	mpp["abc"] = 1
	fmt.Println(mpp)
	fmt.Println(mp["bcd"])
}
