/**

loops workshop3.go
  
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

package loops

import "fmt"

// 220 ile 284 arkadaş sayılardır
// 10 ve 65 arkadaş sayı değildir

func Demo5() {
	number1 := 220
	number2 := 284

	total1 := 0
	total2 := 0

	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			total1 = total1 + i
		}
	}

	for i := 1; i < number2; i++ {
		if number2%i == 0 {
			total2 = total2 + i
		}
	}

	if total1 == number2 && total2 == number1 {
		fmt.Printf("%v ve %v Arkadaş sayılardır\n", number1, number2)
	} else {
		fmt.Printf("%v ve %v Arkadaş sayılar değillerdir\n", number1, number2)
	}
}
