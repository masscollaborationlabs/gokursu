package functions

import "fmt"

func Addition(number1 int, number2 int) int{
	var total = number1 + number2
	return total
}

func SayHello(userName string) {
	fmt.Println("Hello There !", userName)
}
