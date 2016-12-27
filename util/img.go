package util

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"

	"github.com/chai2010/webp"
	"golang.org/x/image/tiff"
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

// EncodeImage ...
func EncodeImage(writer io.Writer, img image.Image, format string, quality float32) error {
	var err error
	switch format {
	case "image/png":
		err = png.Encode(writer, img)
		if err != nil {
			return err
		}
	case "image/jpeg":
		fallthrough
	case "image/jpg":
		err = jpeg.Encode(writer, img, &jpeg.Options{
			Quality: int(quality),
		})
		if err != nil {
			return err
		}
	case "image/webp":
		err = webp.Encode(writer, img, &webp.Options{
			Quality: quality,
		})
		if err != nil {
			return err
		}
	case "image/tiff":
		fallthrough
	case "image/x-tiff":
		err = tiff.Encode(writer, img, &tiff.Options{
			Compression: tiff.Deflate,
		})
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown mime-type: %s", format)
	}
	return nil
}
