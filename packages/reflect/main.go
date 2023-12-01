package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Example 1: Using TypeOf to get the dynamic type
	var num int = 0
	typeOfNum := reflect.TypeOf(num)
	fmt.Printf("TypeOf num: %v\n", typeOfNum)
	// zeroValueOfNum := reflect.Zero(typeOfNum)
	// fmt.Println("Zero Value of Num:", zeroValueOfNum)
	// fmt.Println("Is zero Value of zeroValueOfNum is zero?", zeroValueOfNum.IsZero())

	// Example 2: Using ValueOf to get the runtime value
	valueOfNum := reflect.ValueOf(num).IsZero()
	fmt.Printf("TypeOf valueOfNum: %v\n", valueOfNum)
	// zeroValueOfNum := reflect.Zero(typeOfNum)
	// fmt.Println("Zero Value of Num:", zeroValueOfNum)
	// fmt.Println("Is zero Value of zeroValueOfNum is zero?", zeroValueOfNum.IsZero())

	// Example 3: Using Zero to get a zero value for a type
	zeroValue := reflect.Zero(typeOfNum)
	fmt.Printf("Zero Value: %v\n", zeroValue.Interface())

	// or
	// fmt.Println("Zero Value: ", zeroValue)

	// Difference b/w Printf() and Println()
	// Printf, you may need to explicitly convert values to interface{} using zeroValue.Interface()
	// Println internally handles the conversion of values to strings
}
