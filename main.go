package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func crop(fp string) {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	imgBounds := img.Bounds()

	originR := color.RGBAModel.Convert(img.At(imgBounds.Min.X, imgBounds.Min.Y)).(color.RGBA).R
	leftBorder := 0
	for l := imgBounds.Min.X; l < 100; l++ {
		c := color.RGBAModel.Convert(img.At(l, imgBounds.Min.Y))
		cR := c.(color.RGBA).R
		if cR < originR-5 {
			leftBorder = l
			fmt.Println("left border", leftBorder)
			break
		}
	}

	rightBorder := 0
	for r := 1; r < 100; r++ {
		co := color.RGBAModel.Convert(img.At(imgBounds.Max.X-r, imgBounds.Min.Y))
		coR := co.(color.RGBA).R
		if coR < originR-5 {
			rightBorder = r - 1
			fmt.Println("right border", rightBorder)
			break
		}
	}

	subimg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(leftBorder, 0, imgBounds.Max.X-rightBorder, imgBounds.Max.Y))

	output, outputErr := os.Create("cropped.png")
	if outputErr != nil {
		log.Fatal(outputErr)
	}
	defer output.Close()
	png.Encode(output, subimg)
}

func main() {
	crop("./body.jpg")
}
