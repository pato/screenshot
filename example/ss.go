package main

import (
	"github.com/pato/screenshot"
	"image/png"
	"os"
)

func main() {
	xconn, err := screenshot.Setup()
	if err != nil {
		panic(err)
	}
	img, err := screenshot.CaptureScreen(xconn)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./ss.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}
