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
	Buffer         []Task
}

func New(maxTask, maxWorker int, Name string) Pool {
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
		log.Printf("[Error][%v][AddTask] Queue was full, adding to buffer", p.Name)
		p.Buffer = append(p.Buffer, task)
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
	//go p.ReleaseBuffer()
}

func (p Pool) ReleaseBuffer() {
	for {
		task := p.Buffer[0]
		p.TaskQueue <- task
		p.Buffer = p.Buffer[1:]
	}
}
