package defer_statement

import "fmt"

func A()  {
	fmt.Println("A function worked")
}

func C()  {
	fmt.Println("C function worked")
}

func D()  {
	fmt.Println("D function worked")
}

func B()  {
	defer A()
	defer C()
	defer D()
	fmt.Println("B function worked")
}


