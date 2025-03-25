package main

import "fmt"

func main() {

	var text string = "Hello Golang World!"

	fmt.Println(text)

	var vat int = 12

	fmt.Println(vat)

	var vat_float float64 = 12.2321

	fmt.Println(vat_float)

	vat_another := 25
	fmt.Println(vat_another)
	fmt.Printf("data type : %T\n", vat_another)

	var the_statement bool = false

	var text1 string = "Engin"
	var text2 string = "Ahmet"

	the_statement = text1 == text2

	fmt.Println(the_statement)
}
