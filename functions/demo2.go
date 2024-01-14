package functions

func MathOps(number1 int, number2 int) (int, int, int, float32){
	add := number1 + number2
	subtract := number1 - number2
	multiply := number1 * number2
	divide := float32(number1 / number2)

	return add, subtract, multiply, divide
}