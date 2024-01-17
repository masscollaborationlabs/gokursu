package examplerange

import "fmt"

func Demo3()  {
	my_dictionary:= map[string]string{"book":"kitap", "table":"masa"}
	
	for k,v := range my_dictionary {
		fmt.Println(k)
		fmt.Println(v)
	}
}