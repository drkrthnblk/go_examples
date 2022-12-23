package main

import (
	"os"
	"fmt"
	"bytes"
	"image/jpeg"
	"encoding/base64"
	"github.com/nfnt/resize"
)

func main() {
	errDir := CreateDir("uploads")
	if errDir != nil {
		panic(errDir)
	}
	// encodedData := ""

	// ImageProc(encodedData)

	errDir = DeleteDir("uploads")
	if errDir != nil {
		panic(errDir)
	}
}

func CreateDir(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

func DeleteDir(dirname string) error {
	err := os.Remove(dirname)
    if err != nil {
        return err
    }
	return nil
}

func ImageProc(encodedData string) error {
	data, err := base64.StdEncoding.Decodestring(encodedData)
	if err != nil {
		return err
	}

	original_image, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	newImage := resize.Resize(135, 115, original_image, resize.Lanczos3)

	buf := new(bytes.Buffer)
	// otions -> 0 - 100, low to high image quality
	err = jpeg.Encode(buf, newImage, &jpeg.Options(75))
	imageBit := buf.Bytes()

	photoBase64 := base64.StdEncoding.EncodeToString([]byte(imageBit))
	fmt.Println(photoBase64)
}