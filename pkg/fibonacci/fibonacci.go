package fibonacci

import (
	"math"
)

const Limit = math.MaxInt / 2

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
	fib = fibonacciBelow(Limit)

	// due to properties of summing up even and odds numbers, only exactly every 3 numbers will be even
	sum := 0
	for i := 1; i < len(fib); i += 3 {
		sum += fib[i]
	}

	return sum
}

func fibonacciBelowChannel(limit int, fib chan int) {
	var el1, el2 = 1, 2
	fib <- el2

	for true {
		el := el1 + el2
		if el >= limit {
			break
		}
		if el % 2 == 0 {
			fib <- el
		}
		el1 = el2
		el2 = el
	}
	close(fib)
}

func SumUpChannel() int {
	var fib = make(chan int)

	go fibonacciBelowChannel(Limit, fib)

	sum := 0
	for item := range fib {
		sum += item
	}

	return sum
}
