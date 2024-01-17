package examplerange

import "fmt"

func Demo1()  {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
}