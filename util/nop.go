package util

// NopWriter is /dev/null
type nopWriter struct{}

// DevNull is a special device.
var DevNull = nopWriter{}

// Write makes nothing
func (nopWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func (nopWriter) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}
