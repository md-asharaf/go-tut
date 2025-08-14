package main

import "fmt" // Import the fmt package for formatted I/O

func main() {
	var integer int16 = 5678                               // Initialize a 16-bit integer
	integer++                                              // Increment the integer
	print(integer)                                         // Print the integer using fmt.Println
	print(divide(int(integer), 2))                         // Divide the integer by 2 and print the result
	booleanValue := true                                   // Initialize a boolean variable
	print(booleanValue)                                    // Print the boolean value
	name := "Asharaf"                                      // Initialize a string variable
	print(name)                                            // Print the string variable
	var pi float32                                         // Declare a float32 variable
	print(pi)                                              // Print the default value of pi (0)
	pi = 3.14                                              // Assign a value to pi
	print(pi)                                              // Print the value of pi
	var intArray []int                                     // Declare an integer slice
	print(intArray)                                        // Print the default value of intArray (nil slice)
	intArray = []int{1, 2, 3, 4, 5}                        // Assign values to intArray
	print(intArray)                                        // Print the intArray
	intSlice := append(intArray, 6, 7, 8)                  // Append values to intArray and assign to intSlice
	floatSlice := append([]float64{1}, []float64{3.14}...) // Create and append to a float64 slice
	print(floatSlice)                                      // Print the floatSlice
	print(intSlice)                                        // Print the intSlice
	stringSlice := make([]string, 4, 5)                    // Create a string slice with length 4 and capacity 5
	print(fmt.Sprintf("%T", stringSlice[3]))               // Print the type of the 4th element in stringSlice
	print(len(stringSlice))                                // Print the length of stringSlice
	print(cap(stringSlice))                                // Print the capacity of stringSlice
	var ages = map[string]uint8{"Asharaf": 23, "Bob": 30}
	ages["Charlie"] = 32 // Add a new entry to the map
	print(ages)          // Print the map with names and ages
}

func divide(num int, den int) int {
	return num / den // Return the result of integer division
}

func print(result any) {
	fmt.Println(result) // Print any type of result using fmt.Println
}
