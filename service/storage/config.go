package storage

type Config struct{
	MaxWorkerPoolTask int
	MaxWorkers int
}

func NewDefaultConfig()Config{
	return Config{
		MaxWorkers: 1,
		MaxWorkerPoolTask: 10,
	}
}