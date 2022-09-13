package main

import (
	"github.com/wernerdweight/go-playground/pkg/helpers"
	"github.com/wernerdweight/go-playground/pkg/image_processing"
	"log"
	"runtime"
	"time"
)

func main() {
	log.Printf("max procs: %d", runtime.NumCPU())
	defer helpers.TimeTrack(time.Now(), "image_processing.ProcessChannel")
	image_processing.ProcessChannel("tests/image_processing/images.csv")
}
