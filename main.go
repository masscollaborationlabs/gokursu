package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/functions"
	"golesson/loops"
	"golesson/maps"
	for_range "golesson/range"
	"golesson/slices"
	"golesson/variables"
)

func main() {
	variables.Demo1()
	fmt.Print()
	conditionals.Demo3()
	loops.Demo5()
	arrays.Demo4()
	slices.Demo2()
	functions.SayHello("Mert Gör")
	functions.Addition(2, 6)
	var total = functions.Addition(3, 8)
	fmt.Println(total)
	fmt.Println(total * 10)

	result1, result2, result3, result4 := functions.MathOps(10, 2)
	fmt.Println("Add :", result1)
	fmt.Println("Subtract :", result2)
	fmt.Println("Multiply : ", result3)
	fmt.Println("Divide : ", result4)

	var result = functions.AddVariadic(1, 3, 4, 5, 3)
	fmt.Println(result)

	fmt.Println(functions.AddVariadic(1, 3, 4))

	numbers_main := []int{1, 2, 3, 4}
	fmt.Println(functions.AddVariadic(numbers_main...))

	maps.Demo1()
	for_range.Demo1()
}
