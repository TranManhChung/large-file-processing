package queue

import (
	"log"
	"sync"
)

var (
	queue msgQueue
	once  sync.Once
)

const Size = 1000

type msgQueue struct {
	queue  chan string
	buffer []string
}

func GetQueue() msgQueue {
	once.Do(func() {
		queue = msgQueue{
			queue: make(chan string, Size),
		}
		//go queue.Release()
	})
	return queue
}

func (q msgQueue) Release() {
	for {
		msg := queue.buffer[0]
		queue.queue <- msg
		queue.buffer = queue.buffer[1:]
	}
}

func (q msgQueue) Publish(msg string) {
	select {
	case queue.queue <- msg:
	default:
		log.Printf("[Queue][Publish] Queue was full, add to buffer")
		queue.buffer = append(queue.buffer, msg)
	}

}

func (q msgQueue) Subscribe() string {
	return <-queue.queue
}
