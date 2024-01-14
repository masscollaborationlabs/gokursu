package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/functions"
	"golesson/loops"
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
	functions.Addition(2,6)
	var total = functions.Addition(3,8)
	fmt.Println(total)
	fmt.Println(total * 10)
	
	result1, result2, result3, result4 := functions.MathOps(10,2)
	fmt.Println("Add :", result1)
	fmt.Println("Subtract :", result2)
	fmt.Println("Multiply : ", result3)
	fmt.Println("Divide : ", result4)
}
