package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	c := make(chan struct{}, 1000)
	d := make(chan struct{}, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			c <- struct{}{}
		}
		for i := 0; i < 1000; i++ {
			d <- struct{}{}
		}
	}()
	for i := 0; i < 1000; i++ {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-c:
			fmt.Println("^_^")
		case <-d:
			fmt.Println("^0^")
		default:
			fmt.Println("!!!")
		}
	}
}
