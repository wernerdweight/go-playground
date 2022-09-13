package image_processing

import (
	"os"
)

func ProcessChannel(filepath string) bool {
	reader, _ := os.Open(filepath)
	defer reader.Close()

	items := loadFromFile(filepath)
	itemCount := len(items)

	var loadChannel = make(chan ImageItem, itemCount)
	var floodFillChannel = make(chan ImageItem, itemCount)
	var cropChannel = make(chan ImageItem, itemCount)
	var generateHistogramChannel = make(chan ImageItem, itemCount)
	var saveImageChannel = make(chan ImageItem, itemCount)
	var saveHistogramChannel = make(chan ImageItem, itemCount)
	var syncChannel = make(chan bool, 2*itemCount)

	go func() {
		for item := range loadChannel {
			loadImage(&item)
			floodFillChannel <- item
		}
		close(floodFillChannel)
	}()

	go func() {
		for item := range floodFillChannel {
			floodFill(&item)
			cropChannel <- item
		}
		close(cropChannel)
	}()

	go func() {
		for item := range cropChannel {
			cropToSquare(&item)
			generateHistogramChannel <- item
			saveImageChannel <- item
		}
		close(generateHistogramChannel)
		close(saveImageChannel)
	}()

	go func() {
		for item := range generateHistogramChannel {
			generateHistogram(&item)
			saveHistogramChannel <- item
		}
		close(saveHistogramChannel)
	}()

	go func() {
		for item := range saveImageChannel {
			saveImage(&item)
			syncChannel <- true
		}
	}()

	go func() {
		for item := range saveHistogramChannel {
			saveHistogram(&item)
			syncChannel <- true
		}
	}()

	for _, item := range items {
		loadChannel <- item
	}
	close(loadChannel)

	for i := 0; i < itemCount; i++ {
		<-syncChannel
	}

	return true
}
