package util

import (
	_ "image/png"
	"testing"
)

func TestLoadImage(t *testing.T) {
	bytes, img, err := LoadImage("../test-images/momiji.png")
	if err != nil {
		t.Fatalf("Failed to load image: %v", err)
	}
	if img == nil {
		t.Fatalf("Failed to load image: nil")
	}
	if len(bytes) <= 0 {
		t.Fatalf("Failed to load image: 0 bytes")
	}
}

func TestLoadImageFromNil(t *testing.T) {
	_, _, err := LoadImage("")
	if err == nil {
		t.Fatalf("Why can we load image from \"\"?")
	}
}
