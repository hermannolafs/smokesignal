package smokesignal

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	_ Server = mockServer{}

	mockPort = "4567"
)

type mockServer struct {
	server *http.Server
}

func NewMockServer() mockServer {
	return mockServer{
		server: &http.Server{
			Addr:    ":" + mockPort,
			Handler: Routes(),
		},
	}
}

func (mock mockServer) Run(quit chan os.Signal) {
	log.Println("Mock server received Run instruction")
	time.Sleep(2 * time.Second)
	if err := mock.server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (mock mockServer) Stop(ctx context.Context) error {
	log.Println("Mock server received Stop instruction")
	return mock.server.Shutdown(ctx)
}

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerRoot)

	return router
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Mock server received request")
	w.Write([]byte(r.Method))
}

