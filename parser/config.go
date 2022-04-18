package parser

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
			Username:   "postgres",
			Password:   "chungtm",
			Address:    "localhost:5432",
			Database:   "postgres",
			DriverName: "postgres",
			Format:     "postgres://%s:%s@%s/%s?sslmode=disable",
		},
	}
}
