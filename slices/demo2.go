package slices

import (
	"fmt"
)

func Demo2()  {
	cities := []string{"Ankara", "İstanbul", "İzmir"} // same as make 
	fmt.Println(cities)
	cities_copy := make([]string, len(cities))
	fmt.Println(cities_copy)
	copy(cities_copy, cities)
	fmt.Println(cities_copy)
	cities = append(cities, "Adana")
	fmt.Println(cities)
	fmt.Println(cities_copy)
}