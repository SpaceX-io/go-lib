package perf

import (
	"fmt"
	"testing"
)

func TestStartPprof(t *testing.T) {
	StartPprof([]string{"10.222.64.67:8081"})

	var i1, i2 int
	var c1, c2, c3 chan int
	select {
	case i1 = <-c1:
		fmt.Printf("received %d from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %d to c2\n", i2)
	case i3, ok := <-c3: // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received %d from c3\n", i3)
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
