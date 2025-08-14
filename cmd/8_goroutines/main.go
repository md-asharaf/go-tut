package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var g = sync.WaitGroup{}
var data = [5]int{3, 2, 7, 4, 5}
var results []int

func main() {
	for i := range data {
		g.Add(1)
		go doSomeOperation(i)
	}
	g.Wait()
}

func doSomeOperation(i int) {
	var random = rand.Float32() * 1000

	time.Sleep(time.Duration(random) * time.Millisecond)
	write(data[i])
	read()
	g.Done()
}

func write(value int) {
	m.Lock()
	results = append(results, value)
	m.Unlock()
}

func read() {
	m.RLock()
	fmt.Println("the results are", results)
	m.RUnlock()
}
