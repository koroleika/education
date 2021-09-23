package main

import (
	"fmt"
	"math/big"
	"time"
)

func factCycleFor(n int, done chan bool) {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		select {
		case <-done:
			return
		default:
			result.Mul(result, big.NewInt(int64(i)))
		}
	}
	fmt.Println(result)
}
// контекст
func main() {
	var n int
	fmt.Print("Input digit: ")
	fmt.Scan(&n)

	done := make(chan bool, 1) // структура

	go factCycleFor(n,done)

	to := time.After(1 * time.Second)
	defer close(done)
	for {
		select {
		case <-to:
			fmt.Println("Время истекло")
			done <- true
			return
		}
	}
}