package slices

import "fmt"

func Demo1()  {
	names := make([]string, 3)

	fmt.Println(names)
	names[0]="Engin"
	names[1]="Derin"
	names[2]="Salih"

	fmt.Println(names)
}