package error_handling

import "fmt"

type borderException struct{
	parameter int	
	message string
}

func (b *borderException) Error()  string{
	return fmt.Sprintf("%d ---- %s", b.parameter, b.message)
}

func GuessIt2(guess int)  (string, error){
	if guess < 1 || guess > 100{
		return "", &borderException{guess, "Out of border"}
	}

	return "You did it", nil
}

