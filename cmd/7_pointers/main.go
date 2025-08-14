package main

import "fmt"

func main() {
	var value *int = new(int)
	*value = 42
	fmt.Printf("the value of value is %v  %d\n", *value, *value)

	i := 2
	value = &i
	*value = 100
	fmt.Println(*value, i)
	i = 78
	fmt.Println(*value, i)

	for j, value := range []int{1, 2, 3} {
		fmt.Println(j, value)
	}
}
