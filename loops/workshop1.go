package loops

import "fmt"

func Demo3() {
	number_in_my_mind := 80
	guess_number := 0

	fmt.Println("Guess a number: ")
	fmt.Scanln(&guess_number)
	fmt.Println(guess_number)

	if guess_number == number_in_my_mind {
		fmt.Println("correct !")
	} else if guess_number < number_in_my_mind {
		fmt.Println("lower")
	} else {
		fmt.Println("bigger")
	}
}
