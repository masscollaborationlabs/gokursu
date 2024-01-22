package channels

func EvenNumber(EvenNumberCn chan int)  {
	total := 0
	for i := 0; i <= 10; i+=2 {
		total = total + i
	}
	EvenNumberCn <- total
}

func OddNumber(OddNumberCn chan int)  {
	total := 0
	for i := 1; i <= 10; i+=2 {
		total = total + 1
	}
	OddNumberCn <- total
}