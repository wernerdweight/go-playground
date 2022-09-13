package image_processing

import (
	"encoding/csv"
	"io"
	"os"
)

func ProcessChannel(filepath string) bool {
	var loadChannel = make(chan ImageItem)
	var floodFillChannel = make(chan ImageItem)
	var cropChannel = make(chan ImageItem)
	var generateHistogramChannel = make(chan ImageItem)
	var saveImageChannel = make(chan ImageItem)
	var saveHistogramChannel = make(chan ImageItem)

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
		}
	}()

	go func() {
		for item := range saveHistogramChannel {
			saveHistogram(&item)
		}
	}()

	reader, _ := os.Open(filepath)
	defer reader.Close()

	rootDirectory := getRootDirectory(filepath)
	csvReader := csv.NewReader(reader)
	csvReader.Read() // skip first line
	for {
		line, err := csvReader.Read()
		if io.EOF == err {
			break
		}
		imageItem := imageItemFromCSV(line, rootDirectory)
		loadChannel <- imageItem
	}

	close(loadChannel)

	return true
}
