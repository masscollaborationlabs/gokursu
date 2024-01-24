package defer_statement

import "fmt"

func Demo2(number int)  string{

	defer DeferFunc()

	if number % 2 == 0{
		return "Even number"
	}

	if number %2 != 0 {
		return "Odd Number"
	}

	return "Not defined"
}

func Test()  {
	result := Demo2(10)
	fmt.Println("Result", result)
}

func DeferFunc()  {
	fmt.Println("End !")
}
