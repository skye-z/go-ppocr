package common

import (
	"log"
	"os"

	"gocv.io/x/gocv"
)

func ReadImage(image_path string) gocv.Mat {
	println("read image")
	img := gocv.IMRead(image_path, gocv.IMReadColor)
	if img.Empty() {
		log.Printf("Could not read image %s\n", image_path)
		os.Exit(1)
	}
	return img
}
