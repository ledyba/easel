package util

import (
	"bytes"
	"image"
	"io/ioutil"
	"os"
)

// LoadImage ...
func LoadImage(fname string) ([]byte, image.Image, error) {
	var err error
	freader, err := os.Open(fname)
	if err != nil {
		return nil, nil, err
	}
	defer freader.Close()
	buff, err := ioutil.ReadAll(freader)
	if err != nil {
		return nil, nil, err
	}
	reader := bytes.NewReader(buff)
	src, _, err := image.Decode(reader)
	if err != nil {
		return nil, nil, err
	}
	return buff, src, nil

}
