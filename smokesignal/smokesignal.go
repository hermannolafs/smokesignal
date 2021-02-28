package smokesignal

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	assertPortIsNotUsedTimeout     = 2 * time.Second
	assertPortIsUsedDefaultTimeout = 10 * time.Second
)

// Implement your server with this interface
type Server interface {
	Run(quit chan os.Signal)
	Stop(ctx context.Context) error
}

type SmokeSignal struct {
	t      *testing.T
	server Server
}

func ApiTest(t *testing.T, server Server) SmokeSignal {
	t.Helper()

	return SmokeSignal{
		t:      t,
		server: server,
	}
}

func (smoke SmokeSignal) AssertPortIsNotUsed(port int) SmokeSignal {
	smoke.t.Helper()

	useState, err := CheckIfPortIsInUse(
		port,
		assertPortIsNotUsedTimeout,
		smoke.t,
	)

	if useState == notUsed && err != nil {
		smoke.t.Fatal(err)
	}

	return smoke
}

func (smoke SmokeSignal) LaunchServer() SmokeSignal {
	smoke.t.Helper()

	dummyChannel := make(chan os.Signal, 1)
	go smoke.server.Run(dummyChannel)

	return smoke
}

func (smoke SmokeSignal) AssertPortIsUsed(port int) SmokeSignal {
	smoke.t.Helper()

	useState, err := CheckIfPortIsInUse(
		port,
		assertPortIsUsedDefaultTimeout,
		smoke.t,
	)

	if useState == used && err != nil {
		smoke.t.Fatal(err)
	}

	return smoke
}

func (smoke SmokeSignal) SendGETRequestToPath(path string) SmokeSignal {
	smoke.t.Helper()
	smoke.t.Logf("Sending get request")
	response, err := http.Get(path)
	if err != nil {
		smoke.t.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)

	smoke.t.Logf("Got response: " + string(body))

	return smoke
}

func (smoke SmokeSignal) StopServer() SmokeSignal {
	smoke.t.Helper()
	smoke.t.Log("Calling Server.Stop(ctx)")
	if err := smoke.server.Stop(context.Background()); err != nil {
		smoke.t.Error(err)
	}

	return smoke
}

func (smoke SmokeSignal) EndTest() {
	smoke.t.Log("Finito")
}