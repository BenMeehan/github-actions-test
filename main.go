package calculator

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the result of dividing two integers.
// It returns -1 if the divisor is zero.
func Divide(a, b int) int {
	if b == 0 {
		return -1
	}
	return a / b
}
