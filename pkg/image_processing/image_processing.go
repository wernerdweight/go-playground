package image_processing

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
