package conditionals

import "fmt"

func Demo1() {
	var account float64 = 100
	var money float64 = 10

	var state bool

	state = money > account

	if state {
		fmt.Print("not enough money ")
	}

	if money <= account {
		fmt.Println("your money is getting ready")
		account = account - money
		fmt.Printf("account is : %v\n", account)
	}
	// fmt.Println("done. money in account : " + fmt.Sprintf("%v", account))
	// fmt.Printf("done. money in account : %v", account)
}
