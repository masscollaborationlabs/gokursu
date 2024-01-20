package structs

import "fmt"

func Demo1()  {
	fmt.Println(product{name:"Computer", unitPrice: 5000, brand: "XYZ", discountRate: 20})
}

type product struct {
	name string
	unitPrice float64
	brand string
	discountRate int
}

