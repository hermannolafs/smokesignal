package smokesignal

import (
	"context"
	"os"
	"testing"
)

func Test_APITest_AssertPortIsNotUsed(t *testing.T) {
	totallyFakeServer := NewMockServer()

	ApiTest(t, totallyFakeServer).
		AssertPortIsNotUsed(mockPort)
}

func Test_APITest_AssertPortIsUsed(t *testing.T) {
	totallyFakeServer := NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	ApiTest(t, totallyFakeServer).
		AssertPortIsUsed(mockPort)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_LaunchServer_AssertPortIsUsed(t *testing.T) {
	totallyFakeServer := NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	ApiTest(t, totallyFakeServer).
		AssertPortIsUsed(mockPort)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_AssertPortIsUsed_End(t *testing.T) {
	totallyFakeServer := NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	ApiTest(t, totallyFakeServer).
		AssertPortIsUsed(mockPort)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_AssertPortIsUsed_SendGETRequestToPath(t *testing.T) {
	totallyFakeServer := NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	ApiTest(t, totallyFakeServer).
		SendGETRequestToPath("http://localhost:" + mockPort)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_Full(t *testing.T) {
	totallyFakeServer := NewMockServer()

	ApiTest(t, totallyFakeServer).
		AssertPortIsNotUsed(mockPort).
		LaunchServer().
		AssertPortIsUsed(mockPort).
		SendGETRequestToPath("http://localhost:" + mockPort).
		StopServer().
		AssertPortIsNotUsed(mockPort).
		EndTest()
}
