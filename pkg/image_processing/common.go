package image_processing

import (
	"encoding/csv"
	"fmt"
	"github.com/anthonynsimon/bild/histogram"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/paint"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"image"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func getRootDirectory(filepath string) string {
	rootDirectoryParts := strings.Split(filepath, "/")
	rootDirectory := strings.Join(rootDirectoryParts[0:len(rootDirectoryParts)-1], "/")
	return rootDirectory
}

func imageItemFromCSV(line []string, rootDirectory string) ImageItem {
	filepathParts := strings.Split(line[0], "/")
	fullFilename := filepathParts[len(filepathParts)-1]
	filename := strings.Split(fullFilename, ".")[0]

	pointToFloodFill := strings.Split(line[2], ":")
	pointToFloodFillX, _ := strconv.Atoi(pointToFloodFill[0])
	pointToFloodFillY, _ := strconv.Atoi(pointToFloodFill[1])

	colorToFloodFill := strings.Split(line[3], ":")
	colorToFloodFillR, _ := strconv.Atoi(colorToFloodFill[0])
	colorToFloodFillG, _ := strconv.Atoi(colorToFloodFill[1])
	colorToFloodFillB, _ := strconv.Atoi(colorToFloodFill[2])
	colorToFloodFillA, _ := strconv.Atoi(colorToFloodFill[3])

	colorTolerance, _ := strconv.Atoi(line[4])
	quality, _ := strconv.Atoi(line[5])

	imageItem := ImageItem{
		Path:             fmt.Sprintf("%s/%s", rootDirectory, line[0]),
		Filename:         filename,
		OutputDirectory:  fmt.Sprintf("%s/%s", rootDirectory, line[1]),
		ImageData:        nil,
		PointToFloodFill: image.Point{X: pointToFloodFillX, Y: pointToFloodFillY},
		ColorToFloodFill: color.RGBA{
			R: uint8(colorToFloodFillR),
			G: uint8(colorToFloodFillG),
			B: uint8(colorToFloodFillB),
			A: uint8(colorToFloodFillA),
		},
		ColorTolerance: uint8(colorTolerance),
		Quality:        quality,
	}
	return imageItem
}

func loadImage(imageItem *ImageItem) {
	img, err := imgio.Open(imageItem.Path)
	if nil != err {
		panic(err)
	}
	imageItem.ImageData = img
	log.Printf("image #%s loaded", imageItem.Filename)
}

func cropToSquare(imageItem *ImageItem) {
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, err := analyzer.FindBestCrop(imageItem.ImageData, 1, 1)
	if nil != err {
		panic(err)
	}
	type SubImager interface {
		SubImage(rectangle image.Rectangle) image.Image
	}
	imageItem.ImageData = imageItem.ImageData.(SubImager).SubImage(topCrop)
	log.Printf("image #%s cropped", imageItem.Filename)
}

func floodFill(imageItem *ImageItem) {
	result := paint.FloodFill(imageItem.ImageData, imageItem.PointToFloodFill, imageItem.ColorToFloodFill, imageItem.ColorTolerance)
	imageItem.ImageData = result
	log.Printf("image #%s filled with color %v", imageItem.Filename, imageItem.ColorToFloodFill)
}

func generateHistogram(imageItem *ImageItem) {
	imageItem.Histogram = histogram.NewRGBAHistogram(imageItem.ImageData).Image()
	log.Printf("hostogram for image #%s generated", imageItem.Filename)
}

func saveImage(imageItem *ImageItem) {
	err := imgio.Save(fmt.Sprintf("%s/%s.jpg", imageItem.OutputDirectory, imageItem.Filename), imageItem.ImageData, imgio.JPEGEncoder(imageItem.Quality))
	if nil != err {
		panic(err)
	}
	log.Printf("image #%s saved", imageItem.Filename)
}

func saveHistogram(imageItem *ImageItem) {
	err := imgio.Save(fmt.Sprintf("%s/%s.histogram.jpg", imageItem.OutputDirectory, imageItem.Filename), imageItem.Histogram, imgio.JPEGEncoder(100))
	if nil != err {
		panic(err)
	}
	log.Printf("histogram for image #%s saved", imageItem.Filename)
}
