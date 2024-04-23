package calculator_test

import (
	"fmt"
	"testing"

	calculator "github.com/benmeehan/go-cal"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("Add(5, 3) = %d; expected %d", result, expected)
	} else {
		fmt.Println("Addition Test Passed")
	}
}

func TestSubtract(t *testing.T) {
	result := calculator.Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	} else {
		fmt.Println("Subtraction Test Passed")
	}
}

func TestMultiply(t *testing.T) {
	result := calculator.Multiply(5, 3)
	expected := 15
	if result != expected {
		t.Errorf("Multiply(5, 3) = %d; expected %d", result, expected)
	} else {
		fmt.Println("Multiplication Test Passed")
	}
}

func TestDivide(t *testing.T) {
	result := calculator.Divide(6, 3)
	expected := 2
	if result != expected {
		t.Errorf("Divide(6, 3) = %d; expected %d", result, expected)
	} else {
		fmt.Println("Division Test Passed")
	}

	result = calculator.Divide(6, 0)
	expected = -1
	if result != expected {
		t.Errorf("Divide(6, 0) = %d; expected %d", result, expected)
	} else {
		fmt.Println("Division by Zero Test Passed")
	}
}
