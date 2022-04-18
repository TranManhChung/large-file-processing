package storage

type Config struct {
	MaxWorkerPoolTask int
	MaxWorkers        int
	WorkerName        string
}

func NewDefaultConfig() Config {
	return Config{
		MaxWorkers:        1,
		MaxWorkerPoolTask: 10,
		WorkerName:        "StorageWorker",
	}
}
