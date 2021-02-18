package smokesignal

import (
	"context"
	"os"
	"testing"
	"time"
)

var (
	assertDefaultTimeout    = 3 * time.Second
	assertPortIsUsedTimeout = 10 * time.Second
)

func Test_CheckIfPortIsInUse(t *testing.T) {
	t.Run("BEFORE; Assert port is unused", assertPortIsNotInUse)

	totallyFakeServer := NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal))

	t.Run("DURING; Assert that port is now in use", assertPortIsUsed)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}

	t.Run("AFTER; Assert port is freed", assertPortIsNotInUse)
}

func assertPortIsNotInUse(t *testing.T) {
	t.Helper()
	assertPort(NOTUSED, assertDefaultTimeout, t)
}

func assertPortIsUsed(t *testing.T) {
	t.Helper()
	assertPort(USED, assertPortIsUsedTimeout, t)
}

func assertPort(want used, timeout time.Duration, t *testing.T) {
	t.Helper()
	portUsed, err := CheckIfPortIsInUse(mockPort, timeout, t)

	if err != nil {
		t.Fatal(err)
	}

	if portUsed != want {
		t.Logf("Got: %s\t | Wanted %s", portUsed, want)
		t.Fail()
	}
}
