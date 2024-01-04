package conditionals

import "fmt"

func Demo2() {
	var account float64 = 1000
	var money float64 = 11000

	if money > account {
		fmt.Println("not enough money")
	} else if money == account {
		fmt.Println("your money is getting ready")
		fmt.Println("warning ! no money in account")
	} else {
		fmt.Println("your money is getting ready")
	}

	// if condition {
	// 	if condition {

	// 	}
	// 	// code
	// }

}
