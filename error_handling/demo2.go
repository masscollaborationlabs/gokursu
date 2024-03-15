package error_handling

import (
	"errors"
	"fmt"
)


func GuessIt(guess int) (string, error){

	number_in_mind := 50 

	if guess < 1 || guess > 100{
		return "", errors.New("Type a number between 1 and 100 ")
	}

	if guess > number_in_mind{
		return "Type a small number", nil
	}

	if guess < number_in_mind{
		return "Type a bigger number", nil
	}

	return "You did it", nil
}

func Demo2()  {
	message, error:= GuessIt(101)
	fmt.Println(message)
	fmt.Println(error)
}