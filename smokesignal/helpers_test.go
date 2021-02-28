package smokesignal

import (
	"context"
	"github.com/hermannolafs/smokesignal/mock"
	"os"
	"testing"
	"time"
)

var (
	assertDefaultTimeout    = 3 * time.Second
	assertPortIsUsedTimeout = 10 * time.Second
)

func Test_CheckIfPortIsInUse(t *testing.T) {
	t.Run("BEFORE; Assert port is unused", assertDefaultPortIsNotInUse)

	totallyFakeServer := mock.NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal))

	t.Run("DURING; Assert that port is now in use", assertDefaultPortIsUsed)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}

	t.Run("AFTER; Assert port is freed", assertDefaultPortIsNotInUse)
}

func assertDefaultPortIsNotInUse(t *testing.T) {
	t.Helper()
	assertDefaultPort(notUsed, assertDefaultTimeout, t)
}

func assertDefaultPortIsUsed(t *testing.T) {
	t.Helper()
	assertDefaultPort(used, assertPortIsUsedTimeout, t)
}

func assertDefaultPort(want PortStatus, timeout time.Duration, t *testing.T) {
	t.Helper()
	portUsed, err := CheckIfPortIsInUse(mock.Port, timeout, t)

	if err != nil {
		t.Fatal(err)
	}

	if portUsed != want {
		t.Logf("Got: %s\t | Wanted %s", portUsed, want)
		t.Fail()
	}
}
