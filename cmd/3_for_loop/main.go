package main

import "fmt"

func main() {
	// Example 1: While-like loop (condition only)
	fmt.Println("=== While-like loop (counting down) ===")
	var count int = 5
	for count > 0 {
		fmt.Println("Count is:", count)
		count--
	}

	// Example 2: Another while-like loop (counting up)
	fmt.Println("\n=== While-like loop (counting up) ===")
	for count < 5 {
		fmt.Println("Incrementing count:", count)
		count++
	}

	// Example 3: Traditional for loop
	fmt.Println("\n=== Traditional for loop ===")
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// Example 4: Range-based loop
	fmt.Println("\n=== Range-based loop ===")
	var array = []uint{10, 20, 30, 40, 50}
	for index, value := range array {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example 5: Infinite loop with break (while true equivalent)
	fmt.Println("\n=== Infinite loop with break ===")
	counter := 0
	for {
		if counter >= 3 {
			break
		}
		fmt.Println("Infinite loop iteration:", counter)
		counter++
	}
}
