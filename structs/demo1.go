package structs

import "fmt"

func Demo1()  {
	fmt.Println(product{"Computer", 5000, "XYZ", 20})
}

type product struct {
	name string
	unitPrice float64
	brand string
	discountRate int
}