package fibonacci_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wernerdweight/go-playground/pkg/fibonacci"
	"testing"
)

func TestSumUp(t *testing.T) {
	assertion := assert.New(t)
	assertion.Equal(3_770_056_902_373_173_214, fibonacci.SumUp())
}

func TestSumUpChannel(t *testing.T) {
	assertion := assert.New(t)
	assertion.Equal(3_770_056_902_373_173_214, fibonacci.SumUpChannel())
}
