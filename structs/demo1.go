package structs

import "fmt"

func Demo1()  {
	fmt.Println(product{"Computer", 5000, "XYZ"})
}

type product struct {
	name string
	unitPrice float64
	brand string
}