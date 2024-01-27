/**

variables demo1.go
  
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

package variables

import "fmt"

func Demo1() {
	/**
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")
	fmt.Print("Hello World !")**/
	// fmt.Print("Dünya")
	// fmt.Print("!")
	var text string = "Hello World !"
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)

	var tax1 int = 15

	fmt.Println(tax1)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	var tax2 float32 = 0.1

	fmt.Println(tax2)
	fmt.Println(100 + 100*tax2)

	tax3 := 25.2
	//tax3 = "Engin"
	fmt.Println(tax3)
	fmt.Printf("data type: %T\n", tax3)

	var statement bool

	var text1 string = "Engin"
	var text2 string = "Ahmet"

	statement = text1 != text2

	fmt.Println(statement)

	fmt.Println(!statement)
}
