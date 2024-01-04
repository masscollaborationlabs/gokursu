package variables

import "fmt"

func Demo1() {
	/**
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")**/
	// fmt.Print("Dünya")
	// fmt.Print("!")
	var text string = "Hello World !"
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)

	var tax1 int = 15

	fmt.Println(tax1)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	var tax2 float32 = 0.1

	fmt.Println(tax2)
	fmt.Println(100 + 100*tax2)

	tax3 := 25.2
	//tax3 = "Engin"
	fmt.Println(tax3)
	fmt.Printf("data type: %T\n", tax3)

	var statement bool

	var text1 string = "Engin"
	var text2 string = "Ahmet"

	statement = text1 != text2

	fmt.Println(statement)

	fmt.Println(!statement)
}
