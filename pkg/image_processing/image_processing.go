package image_processing

import (
	"encoding/csv"
	"io"
	"os"
)

func Process(filepath string) bool {
	images := loadFromFile(filepath)
	for _, image := range images {
		loadImage(&image)
		floodFill(&image)
		cropToSquare(&image)
		generateHistogram(&image)
		saveImage(&image)
		saveHistogram(&image)
	}
	return true
}

func loadFromFile(filepath string) []ImageItem {
	var imageItems []ImageItem
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
		imageItems = append(imageItems, imageItem)
	}
	return imageItems
}
