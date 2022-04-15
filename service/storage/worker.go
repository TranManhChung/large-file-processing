package storage

import (
	"github.com/TranManhChung/large-file-processing/service/common/util"
	"io"
	"log"
)

const MaxLine = 100000

type WorkerPool struct {
	TaskQueue      chan string
	NumberOfWorker int
}

func NewWorkerPool(maxTask, maxWorker int) WorkerPool {
	return WorkerPool{
		TaskQueue:      make(chan string, maxTask),
		NumberOfWorker: maxWorker,
	}
}

func (p WorkerPool) AddTask(task string) {
	select {
	case p.TaskQueue <- task:
	default:
		log.Printf("[Error][AddTask] Queue was full, parse this file with tool, file: %v", task)
	}

}

func (p WorkerPool) Run() {
	for i := 0; i < p.NumberOfWorker; i++ {
		go func(taskQueue <-chan string) {
			for task := range taskQueue {
				if err := util.SplitFile(task, MaxLine);err!=nil&&err!=io.EOF {
					log.Printf("[Error][Worker] Cannot handle this task, detail: %v",err)
				}
			}
		}(p.TaskQueue)
	}
}
