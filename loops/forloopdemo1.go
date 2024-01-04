package loops

import "fmt"

func Demo1() {
	var text string = "Hello World !\n"
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)
	// fmt.Println(text)

	i := 1
	
	for i <= 10 {
		fmt.Println(text)
		i = i + 1
	}

	fmt.Println("Done !")
}
