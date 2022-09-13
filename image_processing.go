package main

import (
	"github.com/wernerdweight/go-playground/pkg/helpers"
	"github.com/wernerdweight/go-playground/pkg/image_processing"
	"time"
)

func main() {
	defer helpers.TimeTrack(time.Now(), "image_processing.Process")
	image_processing.Process("tests/image_processing/images.csv")
}
