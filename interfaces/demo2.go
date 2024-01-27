/**

interfaces demo2.go
  
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

package interfaces

import "fmt"

type Mortgage struct{
	creditPaymentTotal float64
	address string
	rate float64
}

type Car struct{
	creditPaymentTotal float64
	carInfo string
	rate float64
}

type CreditCalculator interface{
	Calculate() float64
}

func (m Mortgage)  Calculate() float64{
	return m.creditPaymentTotal * m.rate / 12
}


func (c Car)  Calculate() float64{
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator)  float64{
	
	paymentTotal:= 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + (credits[i].Calculate())		
	}
	return paymentTotal
}

func Demo2()  {

	credit1:= Mortgage{rate:10, creditPaymentTotal: 100000, address: "Ankara"}
	
	credit2:= Mortgage{rate:12, creditPaymentTotal: 50000, address: "İstanbul"}
	
	credit3:= Car{rate:15, creditPaymentTotal: 60000, carInfo:"Polo"}

	credits:= []CreditCalculator{credit1, credit2, credit3}
	
	total:= CalculateMonthlyPayment(credits)

	fmt.Println("Total payment", total)

}
