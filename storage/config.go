package storage

type Config struct {
	Worker Worker
}

type Worker struct {
	MaxWorkerPoolTask int
	MaxWorkers        int
	WorkerName        string
}

func NewDefaultConfig() Config {
	return Config{
		Worker: Worker{
			MaxWorkers:        1,
			MaxWorkerPoolTask: 10,
			WorkerName:        "StorageWorker",
		},
	}
}
