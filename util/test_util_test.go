package util

import "testing"

func TestTestUtil(t *testing.T) {
	StartupTest()
	defer ShutdownTest()
}
