package storage

import "net/http"

type Service struct {
	WorkerPool WorkerPool
}

func NewService(cfg Config) {
	service:= Service{
		WorkerPool: NewWorkerPool(cfg.MaxWorkerPoolTask, cfg.MaxWorkers),
	}

	service.WorkerPool.Run()

	http.HandleFunc("/upload", service.upload)
}

