package example

import (
	"fmt"
	"testing"

	"github.com/hermannolafs/smokesignal/mock"
	"github.com/hermannolafs/smokesignal/smokesignal"
)

func Test_ExampleOfServerInterface(t *testing.T) {
	exampleServer := mock.NewMockServer()

	smokesignal.ApiTest(t, exampleServer).
		AssertPortIsNotUsed(mock.Port).
		LaunchServer().
		AssertPortIsUsed(mock.Port).
		SendGETRequestToPath(fmt.Sprintf("http://localhost:%d", mock.Port)).
		StopServer().
		AssertPortIsNotUsed(mock.Port).
		EndTest()
}
