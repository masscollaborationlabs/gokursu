package string_functions

//alias
import (

	"fmt"

	s "strings"
)

func Demo2()  {
	name := "Engin"
	fmt.Println(s.HasPrefix(name, "Eng")) // true 
	fmt.Println(s.HasSuffix(name, "in")) // true
	fmt.Println(s.Index(name, "in")) // index starts from 0 / zero

	letters := []string{"e", "n", "g","i","n"}
	fmt.Println(s.Join(letters, "*"))
	result := s.Join(letters, "*")
	fmt.Println(result)

	fmt.Println(s.Replace(result, "*", "+", -1))
	fmt.Println(s.Replace(result, "*", "+", 3))

	// For banking : ID Number : 123123121, Date: 02022021, Amount : 5999

	fmt.Println(s.Split(result, "*"))
	fmt.Println(len(s.Split(result, "*")))
	fmt.Println(s.Split(result, "-"))
	fmt.Println(s.Repeat(result, 5))

}