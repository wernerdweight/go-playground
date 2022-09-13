package image_processing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {
	assertion := assert.New(t)
	assertion.True(Process("../../tests/image_processing/images.csv"))
}

func TestProcessChannel(t *testing.T) {
	assertion := assert.New(t)
	assertion.True(ProcessChannel("../../tests/image_processing/images.csv"))
}
