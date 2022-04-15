package worker

import (
	"io"
	"log"
)

type Task func() error

type Pool struct {
	TaskQueue      chan Task
	NumberOfWorker int
	Name           string
}

func NewWorkerPool(maxTask, maxWorker int, Name string) Pool {
	return Pool{
		TaskQueue:      make(chan Task, maxTask),
		NumberOfWorker: maxWorker,
		Name:           Name,
	}
}

func (p Pool) AddTask(task Task) {
	select {
	case p.TaskQueue <- task:
	default:
		log.Printf("[Error][%v][AddTask] Queue was full", p.Name)
	}

}

func (p Pool) Run() {
	for i := 0; i < p.NumberOfWorker; i++ {
		go func(taskQueue <-chan Task) {
			for task := range taskQueue {
				if err := task(); err != nil && err != io.EOF {
					log.Printf("[Error][Worker] Cannot handle this task, detail: %v", err)
				}
			}
		}(p.TaskQueue)
	}
}
