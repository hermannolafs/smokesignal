package smokesignal

import (
	"context"
	"fmt"
	"github.com/hermannolafs/smokesignal/mock"
	"os"
	"testing"
)

func Test_APITest_AssertPortIsNotUsed(t *testing.T) {
	totallyFakeServer := mock.NewMockServer()

	ApiTest(t, totallyFakeServer).
		AssertPortIsNotUsed(mock.Port)
}

func Test_APITest_AssertPortIsUsed(t *testing.T) {
	totallyFakeServer := mock.NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	ApiTest(t, totallyFakeServer).
		AssertPortIsUsed(mock.Port)

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_LaunchServer(t *testing.T) {
	totallyFakeServer := mock.NewMockServer()

	ApiTest(t, totallyFakeServer).
		LaunchServer()

	portUsed, err := CheckIfPortIsInUse(mock.Port, assertPortIsUsedDefaultTimeout, t)
	if err != nil || portUsed != used {
		t.Fatalf("Got %s | Wanted %s \n Err: %+v", portUsed, used, err)
	}

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_SendGETRequestToPath(t *testing.T) {
	totallyFakeServer := mock.NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	portUsed, err := CheckIfPortIsInUse(mock.Port, assertPortIsUsedDefaultTimeout, t)
	if err != nil || portUsed != used {
		t.Fatalf("Got %s | Wanted %s \n Err: %+v", portUsed, used, err)
	}

	ApiTest(t, totallyFakeServer).
		SendGETRequestToPath(fmt.Sprintf("http://localhost:%d", mock.Port))

	if err := totallyFakeServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func Test_APITest_StopServer(t *testing.T) {
	totallyFakeServer := mock.NewMockServer()
	go totallyFakeServer.Run(make(chan os.Signal, 1))

	portUsed, err := CheckIfPortIsInUse(mock.Port, assertPortIsUsedDefaultTimeout, t)
	if err != nil || portUsed != used {
		t.Fatalf("Got %s | Wanted %s \n Err: %+v", portUsed, used, err)
	}

	ApiTest(t, totallyFakeServer).
		StopServer()

	portUsed, err = CheckIfPortIsInUse(mock.Port, assertPortIsUsedDefaultTimeout, t)
	if err != nil || portUsed != notUsed {
		t.Fatalf("Got %s | Wanted %s \n Err: %+v", portUsed, used, err)
	}
}


func Test_APITest_Full(t *testing.T) {
	totallyFakeServer := mock.NewMockServer()

	ApiTest(t, totallyFakeServer).
		AssertPortIsNotUsed(mock.Port).
		LaunchServer().
		AssertPortIsUsed(mock.Port).
		SendGETRequestToPath(fmt.Sprintf("http://localhost:%d", mock.Port)).
		StopServer().
		AssertPortIsNotUsed(mock.Port).
		EndTest()
}
