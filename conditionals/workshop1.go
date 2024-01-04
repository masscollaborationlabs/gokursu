package conditionals

import "fmt"

func Demo3() {
	// thre int variables defined
	// print screen the max int

	// var a int = 1
	// var b int = 2
	// var c int = 3

	var number1, number2, number3 int = 10, 5, 18

	var max int = number1

	if number2 > max {
		max = number2
	}

	if number3 > max {
		max = number3
	}

	fmt.Printf("max number is: %v\n", max)
}
