package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Understanding Channels Step by Step ===")
	fmt.Println()

	// Step 1: The simplest possible channel example
	simpleChannelExample()

	// Step 2: Why do we need channels?
	whyChannelsExample()

	// Step 3: Practical example
	practicalExample()
}

// Step 1: The absolute basics
func simpleChannelExample() {
	fmt.Println("STEP 1: Basic Channel (like passing a note)")
	fmt.Println("=========================================")

	// Create a "mailbox" that can hold strings
	mailbox := make(chan string)

	// Person A puts a note in the mailbox (in a separate goroutine)
	go func() {
		fmt.Println("Person A: Writing a note...")
		time.Sleep(1 * time.Second)       // Taking time to write
		mailbox <- "Hello from Person A!" // Put note in mailbox
		fmt.Println("Person A: Note placed in mailbox!")
	}()

	// Person B waits and reads the note (main goroutine)
	fmt.Println("Person B: Waiting for a note...")
	note := <-mailbox // Take note from mailbox (BLOCKS until note arrives)
	fmt.Println("Person B: Got the note:", note)
	fmt.Println()
}

// Step 2: Why channels are needed
func whyChannelsExample() {
	fmt.Println("STEP 2: Why we need channels")
	fmt.Println("=============================")

	// Without channels - this would be a problem:
	// var sharedMessage string // This is shared memory - DANGEROUS!

	// Create a channel instead
	messageCh := make(chan string)

	// Worker goroutine
	go func() {
		fmt.Println("Worker: Doing some work...")
		time.Sleep(2 * time.Second)

		// Instead of: sharedMessage = "Work done!" (unsafe)
		// We use channel: (safe)
		messageCh <- "Work completed safely!"
		fmt.Println("Worker: Sent message via channel")
	}()

	// Main goroutine waits for worker
	fmt.Println("Main: Waiting for worker to finish...")
	result := <-messageCh
	fmt.Println("Main: Worker said:", result)
	fmt.Println()
}

// Step 3: A practical example
func practicalExample() {
	fmt.Println("STEP 3: Practical Example - Calculator Service")
	fmt.Println("==============================================")

	// Create channels for sending numbers and receiving results
	numbersCh := make(chan int)
	resultsCh := make(chan int, 3) // Buffered channel for 3 results

	// Start a calculator goroutine (like a service)
	go calculator(numbersCh, resultsCh)

	// Send numbers to calculate
	fmt.Println("Sending numbers: 5, 10, 15")
	numbersCh <- 5
	numbersCh <- 10
	numbersCh <- 15
	close(numbersCh) // Tell calculator no more numbers coming

	// Receive results
	fmt.Println("Receiving results:")
	for result := range resultsCh { // Keep reading until channel closes
		fmt.Printf("Result: %d\n", result)
	}
}

// Calculator function that doubles numbers
func calculator(input <-chan int, output chan<- int) {
	fmt.Println("Calculator: Started, waiting for numbers...")

	for number := range input { // Read each number sent to us
		fmt.Printf("Calculator: Processing %d\n", number)
		result := number * 2 // Double it
		output <- result     // Send result back
	}

	close(output) // Close output channel when done
	fmt.Println("Calculator: Finished!")
}
