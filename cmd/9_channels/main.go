package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Basic channel usage
	basicChannelExample()

	// Example 2: Channels with goroutines
	channelWithGoroutinesExample()

	// Example 3: Buffered channels
	bufferedChannelExample()

	// Example 4: Channel directions (send-only, receive-only)
	channelDirectionsExample()

	// Example 5: Select statement
	selectExample()
}

// Example 1: Basic Channel Operations
func basicChannelExample() {
	fmt.Println("=== Basic Channel Example ===")

	// Create a channel that can send/receive strings
	ch := make(chan string)

	// Start a goroutine that will send data
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!" // Send data to channel
	}()

	// Receive data from channel (this blocks until data arrives)
	message := <-ch
	fmt.Println("Received:", message)
	fmt.Println()
}

// Example 2: Using Channels for Goroutine Communication
func channelWithGoroutinesExample() {
	fmt.Println("=== Channels with Goroutines ===")

	ch := make(chan int)

	// Start 3 goroutines that send numbers
	for i := 1; i <= 3; i++ {
		go func(num int) {
			time.Sleep(time.Duration(num) * time.Second)
			ch <- num * num // Send square of number
		}(i)
	}

	// Receive 3 values
	for i := 0; i < 3; i++ {
		result := <-ch
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println()
}

// Example 3: Buffered Channels
func bufferedChannelExample() {
	fmt.Println("=== Buffered Channels ===")

	// Create a buffered channel (can hold 3 values without blocking)
	ch := make(chan string, 3)

	// Send 3 values (won't block because buffer size is 3)
	ch <- "first"
	ch <- "second"
	ch <- "third"

	// Receive them
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println()
}

// Example 4: Channel Directions
func channelDirectionsExample() {
	fmt.Println("=== Channel Directions ===")

	ch := make(chan string)

	go sender(ch) // Pass channel as send-only
	receiver(ch)  // Pass channel as receive-only
	fmt.Println()
}

// Function that can only send to channel
func sender(ch chan<- string) {
	ch <- "Message from sender function"
}

// Function that can only receive from channel
func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Println("Received in receiver function:", msg)
}

// Example 5: Select Statement (non-blocking operations)
func selectExample() {
	fmt.Println("=== Select Statement ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Two goroutines sending at different times
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Select waits for the first available channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Got from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Got from ch2:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}
