package query

type Config struct {
	Postgresql Postgresql
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