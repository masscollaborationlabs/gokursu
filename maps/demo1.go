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
	fmt.Println("Dictionary", my_dictionary)
	delete(my_dictionary, "book")
	fmt.Println("String number : ", len(my_dictionary))
	fmt.Println("Dictionary", my_dictionary)

	value, is_value_there := my_dictionary["table"]
	fmt.Println(value)
	fmt.Println("Listede olma durumu", is_value_there)
}