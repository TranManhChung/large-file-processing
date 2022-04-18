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
			1,
			10,
			"ParserWorker",
		},
		Postgresql: Postgresql{
			"postgres",
			"chungtm",
			"localhost:5432",
			"postgres",
			"postgres",
			"postgres://%s:%s@%s/%s?sslmode=disable",
		},
	}
}
