package defer_statement

import "fmt"

func A()  {
	fmt.Println("A function worked")
}

func B()  {
	fmt.Println("B function worked")
	A()
}

