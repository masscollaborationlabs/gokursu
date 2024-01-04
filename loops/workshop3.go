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
