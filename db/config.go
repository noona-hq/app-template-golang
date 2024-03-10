package db

type Config struct {
	Connection       string `default:"mongodb://localhost:31060"`
	Name             string `default:"app_template"`
	DirectConnection bool   `default:"true"`
}
