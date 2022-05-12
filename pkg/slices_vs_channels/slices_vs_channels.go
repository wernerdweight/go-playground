package slices_vs_channels

import (
	"math/rand"
)

const (
	IterationLimit = 1_000
	RandLimit = 100_000
)

func randomBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < RandLimit; i++ {
		rand.Read(b)
	}
	return b
}

func randomNumbers(limit int) []int {
	var fib []int

	for i := 0; i < limit; i++ {
		fib = append(fib, len(randomBytes(100)))
	}

	return fib
}

func SumUp() int {
	// sum up even numbers in fibonacci

	var fib []int
	fib = randomNumbers(IterationLimit)

	// due to properties of summing up even and odds numbers, only exactly every 3 numbers will be even
	sum := 0
	for i := 0; i < len(fib); i++ {
		sum += fib[i]
	}

	return sum
}

func randomNumbersChannel(limit int, fib chan int) {
	for i := 0; i < limit; i++ {
		fib <- len(randomBytes(100))
	}
	close(fib)
}

func SumUpChannel() int {
	var fib = make(chan int)

	go randomNumbersChannel(IterationLimit, fib)

	sum := 0
	for item := range fib {
		sum += item
	}

	return sum
}
