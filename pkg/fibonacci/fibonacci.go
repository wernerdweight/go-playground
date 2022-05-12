package fibonacci

import (
	"fmt"
	"math"
)

func fibonacciBelow(limit int) []int {
	var fib []int
	fib = append(fib, 1)
	fib = append(fib, 2)

	var i = 1
	for fib[i] < limit {
		fib = append(fib, fib[i-1]+fib[i])
		i++
	}

	return fib[:len(fib)-1]
}

func SumUp() int {
	// sum up even numbers in fibonacci

	var fib []int
	for i := 0; i < 10_000_000; i++ {
		fib = fibonacciBelow(math.MaxInt / 2)
	}

	// due to properties of summing up even and odds numbers, only exactly every 3 numbers will be even
	sum := 0
	for i := 1; i < len(fib); i += 3 {
		sum += fib[i]
	}

	fmt.Println(sum)
	return sum
}
