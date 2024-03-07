package coderabbit

import (
	"testing"
)

func TestSumUp(t *testing.T) {
	expected := 4613732 // calculated sum of the first 50 Fibonacci numbers
	actual := SumUp()
	if actual != expected {
		t.Errorf("SumUp() = %d; want %d", actual, expected)
	}
}

func TestSumUpChannel(t *testing.T) {
	expected := 4613732 // calculated sum of the first 50 Fibonacci numbers using channels
	actual := SumUpChannel()
	if actual != expected {
		t.Errorf("SumUpChannel() = %d; want %d", actual, expected)
	}
}

func TestRandomBytes(t *testing.T) {
	length := 100
	bytes := randomBytes(length)
	if len(bytes) != length {
		t.Errorf("randomBytes(%d) returned slice of length %d; want %d", length, len(bytes), length)
	}
}

func TestRandomNumbers(t *testing.T) {
	limit := 50
	nums := randomNumbers(limit)
	if len(nums) != limit {
		t.Errorf("randomNumbers(%d) returned slice of length %d; want %d", limit, len(nums), limit)
	}
}

func TestSumUpSlices(t *testing.T) {
	// Provide a known slice to sum up
	fib := []int{1, 2, 3, 4, 5}
	expected := 15 // sum of the provided slice
	actual := SumUpSlices()
	if actual != expected {
		t.Errorf("SumUpSlices() = %d; want %d", actual, expected)
	}
}
