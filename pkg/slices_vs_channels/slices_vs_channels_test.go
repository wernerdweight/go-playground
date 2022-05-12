package slices_vs_channels_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wernerdweight/tmp/pkg/slices_vs_channels"
	"testing"
)

func TestSumUp(t *testing.T) {
	assertion := assert.New(t)
	assertion.Greater(slices_vs_channels.SumUp(), 0)
}

func TestSumUpChannel(t *testing.T) {
	assertion := assert.New(t)
	assertion.Greater(slices_vs_channels.SumUpChannel(), 0)
}
