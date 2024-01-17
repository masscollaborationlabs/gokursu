package examplerange

import "fmt"

// add all odd numbers between between 1 and 10
// for range
func Demo2()  {
	numbers:=[]int{1,2,3,4,5,6,7,8,9,10}
	total:=0

	for _, number:= range numbers{
		 if number%2 !=0{
			total = total + number
		 }
	}
	fmt.Println("Total is : ", total)
}