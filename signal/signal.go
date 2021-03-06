package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ShutdownSignal(shutdownFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Printf("server start success pid:%d\n", os.Getpid())

	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			shutdownFunc()
			return
		default:
			return
		}
	}
}
