package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	img, err := vips.NewImageFromFile("example.jpg")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	imgBytes, _, err := img.ExportJpeg(vips.NewJpegExportParams())
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	ioutil.WriteFile("out.jpg", imgBytes, 0644)
}
