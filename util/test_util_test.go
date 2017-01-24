package util

import "testing"

func TestTestUtil(t *testing.T) {
	StartupTest(t)
	defer ShutdownTest()
}
