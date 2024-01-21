package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/examplerange"
	"golesson/functions"
	"golesson/goroutines"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	"golesson/slices"
	"golesson/structs"
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
	examplerange.Demo1()
	examplerange.Demo2()
	examplerange.Demo3()

	number := 20
	pointers.Demo1(&number)
	fmt.Println("Number in Main go file", number)

	numbers:=[]int{1,2,3}
	pointers.Demo2(numbers)
	fmt.Println("Numbers in Main", numbers[0])
	structs.Demo1()
	structs.Demo2()

	goroutines.EvenNumber()
	goroutines.OddNumber()
}
