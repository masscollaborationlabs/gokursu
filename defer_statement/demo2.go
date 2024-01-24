package defer_statement

import "fmt"

func Demo2(number int)  string{
	if number % 2 == 0{
		return "Even number"
	}

	if number %2 != 0 {
		return "Odd Number"
	}

	return "Not defined"
}

func Test()  {
	result := Demo2(9)
	fmt.Println("Result", result)
}

