package signal

import (
	"flag"
	"fmt"
	"net/http"
	"testing"
	"time"
)

var addr = flag.String("p", "10.222.64.67:8081", "port")

func TestShutdownSignal(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/testApi", testApi)
	s := &http.Server{
		Handler:        mux,
		Addr:           *addr,
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			t.Error(err)
		}
	}()
	ShutdownSignal(func() {
		err := s.Close()
		if err != nil {
			t.Error(err)
		}
		fmt.Println("server quit!")
	})
}

func testApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testApi access!")
}
