package storage

import (
	"github.com/TranManhChung/large-file-processing/service/common/worker"
	"net/http"
)

type Service struct {
	WorkerPool worker.Pool
}

func NewService(cfg Config) {
	service:= Service{
		WorkerPool: worker.NewWorkerPool(cfg.MaxWorkerPoolTask, cfg.MaxWorkers,"StorageWorker"),
	}

	service.WorkerPool.Run()

	http.HandleFunc("/upload", service.upload)
}

