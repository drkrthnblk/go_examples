package main

import (
	"os"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(path+"/main/img/hdimg2.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/jpg":
		base64Encoding += "data:image/jpg;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)
}