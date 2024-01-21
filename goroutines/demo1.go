package goroutines

import (
	"fmt"
	"time"
)

func EvenNumber()  {
	for i := 0; i <= 10; i+=2 {
		fmt.Println("Even Number : ", i)
		time.Sleep(1 * time.Second)
	}
}

func OddNumber()  {
	for i := 1; i <= 10; i+=2 {
		fmt.Println("Odd Number : ", i)
		time.Sleep(1 * time.Second)
	}
}