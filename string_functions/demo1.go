package string_functions

//alias
import (
	"fmt"
	s "strings"
)

// case sensitive
// ascii
func Demo1() {
	name := "Engin"
	fmt.Println(s.Count(name, "g"))
	fmt.Println(s.Contains(name, "g"))
	fmt.Println(s.Index(name, "g"))
	result := s.Index(name, "E")

	if result != -1{
		fmt.Println("E letter exists")
	}else{
		fmt.Println("E letter does not exist")
	}

	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))
}
