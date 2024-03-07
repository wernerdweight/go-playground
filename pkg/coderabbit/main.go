package coderabbit

func init() {
	// This is the init function
}

// SumUp returns the sum of the first 50 fibonacci numbers
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

// SumUpChannel returns the sum of the first 50 fibonacci numbers
func SumUpChannel() int {
	var fib = make(chan int)

	go fibonacciBelowChannel(Limit, fib)

	sum := 0
	for item := range fib {
		sum += item
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
		if el%2 == 0 {
			fib <- el
		}
		el1 = el2
		el2 = el
	}
	close(fib)
}

// Limit is the limit of the fibonacci sequence
const Limit = 4_000_000

// IterationLimit is the limit of the iteration
const IterationLimit = 1_000

// RandLimit is the limit of the random number
const RandLimit = 100_000

// randomBytes returns a random byte array
func randomBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < RandLimit; i++ {
		rand.Read(b)
	}
	return b
}

// randomNumbers returns a random number array
func randomNumbers(limit int) []int {
	var fib []int

	for i := 0; i < limit; i++ {
		fib = append(fib, len(randomBytes(100)))
	}

	return fib
}

// SumUpSlices returns the sum of the random numbers
func SumUpSlices() int {
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

// randomNumbersChannel returns a random number array
func randomNumbersChannel(limit int, fib chan int) {
	for i := 0; i < limit; i++ {
		fib <- len(randomBytes(100))
	}
	close(fib)
}

// SumUpChannelSlices returns the sum of the random numbers
func SumUpChannelSlices() int {
	var fib = make(chan int)

	go randomNumbersChannel(IterationLimit, fib)

	sum := 0
	for item := range fib {
		sum += item
	}

	return sum
}
