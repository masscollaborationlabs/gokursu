/**

main.go

BSD 3-Clause License

Copyright (c) 2024 Mass Collaboration Labs and contributors

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

 **/

package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/channels"
	"golesson/conditionals"
	"golesson/defer_statement"
	"golesson/error_handling"
	"golesson/examplerange"
	"golesson/functions"
	"golesson/goroutines"
	"golesson/interfaces"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	"golesson/restful"
	"golesson/slices"
	"golesson/string_functions"
	"golesson/structs"
	"golesson/variables"
	"time"
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

	numbers := []int{1, 2, 3}
	pointers.Demo2(numbers)
	fmt.Println("Numbers in Main", numbers[0])
	structs.Demo1()
	structs.Demo2()

	go goroutines.EvenNumber()
	go goroutines.OddNumber()
	time.Sleep(5 * time.Second)
	fmt.Println("Main ended")

	EvenNumberCn := make(chan int)
	OddNumberCn := make(chan int)

	go channels.EvenNumber(EvenNumberCn)
	go channels.OddNumber(OddNumberCn)

	EvenNumberTotal, OddNumberTotal := <-EvenNumberCn, <-OddNumberCn

	multiply := EvenNumberTotal * OddNumberTotal

	fmt.Println("Multiply Result : ", multiply)

	interfaces.Demo1()
	interfaces.Demo2()

	defer_statement.B()
	defer_statement.Test()
	defer_statement.Demo3()

	error_handling.Demo1()
	interfaces.Demo3()
	error_handling.Demo2()
	fmt.Println(error_handling.GuessIt2(102))
	string_functions.Demo1()
	string_functions.Demo2()
	restful.Demo1()
	restful.Demo2()
}
