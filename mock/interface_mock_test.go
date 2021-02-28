package mock_test

import (
	"context"
	"fmt"
	"github.com/hermannolafs/smokesignal/mock"
	"github.com/hermannolafs/smokesignal/smokesignal"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	_ smokesignal.Server = mock.Server{}
)

func TestNewMockServer(t *testing.T) {
	firstTestServer := mock.NewMockServer()

	go firstTestServer.Run(make(chan os.Signal, 1))

	time.Sleep(3 * time.Second) // Wait for artificial startup time to elapse

	assertMockServerRespondsWithMethodOnRootPath(t)

	if err := firstTestServer.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func assertMockServerRespondsWithMethodOnRootPath(t *testing.T) {
	t.Helper()
	response := sendGetRequestToRootEndpoint(t)

	body := readBodyFromResponse(t, response)

	assertResponseBodyIsMethod(t, body)
}

func assertResponseBodyIsMethod(t *testing.T, body string) {
	t.Helper()
	if body != http.MethodGet {
		t.Fatal(fmt.Sprintf("Expected %s, Got %s", http.MethodGet, body))
	}
}

func readBodyFromResponse(t *testing.T, response *http.Response) string {
	t.Helper()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(body)
}

func sendGetRequestToRootEndpoint(t *testing.T) *http.Response {
	t.Helper()
	response, err := http.Get(fmt.Sprintf("http://localhost:%d", mock.Port))
	if err != nil {
		t.Fatal(err)
	}
	return response
}