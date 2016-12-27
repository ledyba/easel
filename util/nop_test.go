package util

import "testing"

func TestDevNull(t *testing.T) {
	buff := make([]byte, 100)
	n, err := DevNull.Read(buff)
	if err != nil {
		t.Error("DevNull must succeed reading operation")
	}
	if n != len(buff) {
		t.Error("DevNull must succeed reading operation with any buffer length")
	}

	n, err = DevNull.Write(buff)
	if err != nil {
		t.Error("DevNull must succeed reading operation")
	}
	if n != len(buff) {
		t.Error("DevNull must succeed reading operation with any buffer length")
	}
}
