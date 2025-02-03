package main

import "fmt"

func main() {
	// fmt.Println("Hello world")

	var name string = "your name"
	fmt.Println(name)

	// looping

	// infinite loop
	// for {
	// 	println(1)
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := range 3 {
	// 	fmt.Println(i)
	// }

	var age int = 18

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Under 18.")
	}

	if p := 12; p == 12 {
		println("p is", p)
	}
	// go does not have ternary operator and while loops

	// switch statements

	a := 5
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("It is ", a)
	}
}
