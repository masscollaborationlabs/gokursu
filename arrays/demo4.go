package arrays

import "fmt"

func Demo4()  {
	var numbers [2][3]int

	numbers[0][0] = 1
	numbers[0][1] = 3
	numbers[0][2] = 5
	numbers[1][0] = 2
	numbers[1][1] = 4
	numbers[1][2] = 6

	fmt.Println(numbers[1][1])
}