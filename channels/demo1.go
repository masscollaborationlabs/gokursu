package channels

import (
	"fmt"
	"time"
)

func EvenNumber(EvenNumberCn chan int)  {
	total := 0
	for i := 0; i <= 10; i+=2 {
		total = total + i
		fmt.Println("Even Number is working")
		time.Sleep(1 * time.Second)
	}
	
	EvenNumberCn <- total
}

func OddNumber(OddNumberCn chan int)  {
	total := 0
	for i := 1; i <= 10; i+=2 {
		total = total + 1
		fmt.Println("Odd Number is working")
		time.Sleep(1 * time.Second)
	}
	OddNumberCn <- total
}