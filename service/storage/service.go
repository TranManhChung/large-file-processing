package storage

import (
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/pkg/worker"
	"net/http"
)

type Service struct {
	WorkerPool worker.Pool
}

func New() func() {
	cfg := NewDefaultConfig()
	service := Service{
		WorkerPool: worker.New(cfg.Worker.MaxWorkerPoolTask, cfg.Worker.MaxWorkers, cfg.Worker.WorkerName),
	}

	service.WorkerPool.Run()

	http.HandleFunc("/upload", service.upload)
	go http.ListenAndServe(":8080", nil)

	return func() {
		fmt.Println("Clean up storage, detail: finished")
	}
}
