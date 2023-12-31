package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("missing first name or last name")
	} else {
		return x / y, nil
	}
}
