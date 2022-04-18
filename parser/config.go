package parser

import "os"

type Config struct {
	Worker     Worker
	Postgresql Postgresql
}

type Worker struct {
	MaxWorkerPoolTask int
	MaxWorkers        int
	WorkerName        string
}

type Postgresql struct {
	Username   string
	Password   string
	Address    string
	Database   string
	DriverName string
	Format     string
}

func NewDefaultConfig() Config {

	return Config{
		Worker: Worker{
			MaxWorkers:        1,
			MaxWorkerPoolTask: 10,
			WorkerName:        "ParserWorker",
		},
		Postgresql: Postgresql{
			Username:   os.Getenv("POSTGRES_USER"),
			Password:   os.Getenv("POSTGRES_PASSWORD"),
			Address:    os.Getenv("POSTGRES_ADDRESS"),
			Database:   os.Getenv("POSTGRES_DATABASE"),
			DriverName: os.Getenv("POSTGRES_DRIVER_NAME"),
			Format:     "postgres://%s:%s@%s/%s?sslmode=disable",
		},
	}
}
