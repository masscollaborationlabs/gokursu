package maps

import "fmt"

func Demo1()  {
	my_dictionary := make(map[string]string)
	// key - value
	my_dictionary["table"]="Masa"
	my_dictionary["book"]="kitap"
	my_dictionary["pencil"]="Kalem"
	
	fmt.Println(my_dictionary["book"])
	fmt.Println("String number : ", len(my_dictionary))
}