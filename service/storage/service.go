package storage

import (
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/common/worker"
	"net/http"
)

type Service struct {
	WorkerPool worker.Pool
}

func New() func() {
	cfg := NewDefaultConfig()
	service := Service{
		WorkerPool: worker.New(cfg.MaxWorkerPoolTask, cfg.MaxWorkers, cfg.WorkerName),
	}

	service.WorkerPool.Run()

	http.HandleFunc("/upload", service.upload)
	go http.ListenAndServe(":8080", nil)

	return func() {
		fmt.Println("Clean up storage, detail: finished")
	}
}
