package quering

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
			Username:   "postgres",
			Password:   "chungtm",
			Address:    "localhost:5432",
			Database:   "postgres",
			DriverName: "postgres",
			Format:     "postgres://%s:%s@%s/%s?sslmode=disable",
		},
	}
}