package defer_statement

import "fmt"

func A()  {
	fmt.Println("A function worked")
}

func B()  {
	defer A()
	fmt.Println("B function worked")
}

