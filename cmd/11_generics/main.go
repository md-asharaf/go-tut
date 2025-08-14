package main

import "fmt"

func main() {
	var num1 int = 34
	var num2 int32 = 56
	var num3 int64 = 635
	var num4 float64 = 4352.87
	print(multiplyBy2(num1))
	print(multiplyBy2(num2))
	print(multiplyBy2(num3))
	print(multiplyBy2(num4))
}

// Generic function with type parameter T
func multiplyBy2[T int | int32 | int64 | float64](value T) T {
	return value * 2
}

func print(value any) {
	fmt.Printf("The type of %v is %T\n", value, value)
}
