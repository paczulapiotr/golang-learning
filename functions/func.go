package functions

import (
	"fmt"
)

func Run() {
	varioticParameters("Message")
	returnValueNormal()
	specialReturn()
	returnPointer()
	tupleReturn(10, 23)
}

func varioticParameters(message string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(message, result)
}

func returnValueNormal() int {
	return 10
}

func specialReturn() (result int) {
	result = 10

	return
}

// valueToReturn wont be cleared from stack => its good in go
func returnPointer() *int {
	valueToReturn := 10

	return &valueToReturn
}

func tupleReturn(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}
