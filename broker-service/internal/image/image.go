package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"
	"os"
)

func imageToString64(path string) string {
	img := openImage(path)
	imgBuf := new(bytes.Buffer)

	//TODO: abstract the encoding method

	_ = png.Encode(imgBuf, img)
	return base64.StdEncoding.EncodeToString(imgBuf.Bytes())
}

func openImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
