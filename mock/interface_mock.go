package mock

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	Port = 4567
)

type Server struct {
	server *http.Server
}

func NewMockServer() Server {
	return Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", Port),
			Handler: Routes(),
		},
	}
}

func (mock Server) Run(quit chan os.Signal) {
	_ = quit
	log.Println("Mock server received Run instruction")
	time.Sleep(2 * time.Second) // Artificially introduce a startup time
	if err := mock.server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (mock Server) Stop(ctx context.Context) error {
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
	if _, err := w.Write([]byte(r.Method)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
