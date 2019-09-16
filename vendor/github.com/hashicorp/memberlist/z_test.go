package memberlist

import (
	"log"
	"strings"
	"testing"
)

func testLogger(t testing.TB) *log.Logger {
	return log.New(&testWriter{t}, "test: ", log.LstdFlags)
}

func testLoggerWithName(t testing.TB, name string) *log.Logger {
	return log.New(&testWriter{t}, "test["+name+"]: ", log.LstdFlags)
}

type testWriter struct {
	t testing.TB
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.t.Helper()
	tw.t.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}
