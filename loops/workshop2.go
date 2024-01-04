package loops

import (
	"fmt"
)

func Demo4() {
	number := 0
	fmt.Println("enter a number")
	fmt.Scanln(&number)

	is_prime := true

	for i := 2; i < number; i++ {
		if number%i == 0 {
			is_prime = false
		}
	}
	if is_prime == true {
		fmt.Println("prime")
	} else {
		fmt.Println("not prime")
	}
}
// I will write it again for now this code is written by the lecturer