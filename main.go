package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./ref.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	subimg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(30, 0, img.Bounds().Dx(), img.Bounds().Dy()))

	output, outputErr := os.Create("output.png")
	if outputErr != nil {
		log.Fatal(outputErr)
	}
	defer output.Close()
	png.Encode(output, subimg)
}
