package perf

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func StartPprof(addrs []string) {
	pprofServeMux := http.NewServeMux()
	pprofServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)

	for _, addr := range addrs {
		go func(addr string) {
			if err := http.ListenAndServe(addr, pprofServeMux); err != nil {
				fmt.Printf("perf http.ListenAndServe error: %s %v", addr, err)
				panic(err)
			}
		}(addr)
	}
}
