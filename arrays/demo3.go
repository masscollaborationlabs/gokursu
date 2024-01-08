package arrays

import "fmt"

func Demo3()  {
	numbers := [5]int{20, 30, 50, 10, 2}
	fmt.Println(numbers)

	max_number := numbers[0]



	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max_number {
			max_number = numbers[i]
		}
		fmt.Println("max number : ", max_number)
	}
}