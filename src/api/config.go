package api

import "os"

var config Config

type Config struct {
	DB    DBConfig
	Redis RedisConfig
	Http  HTTPConfig
}
type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

type HTTPConfig struct {
	Port string
}

type RedisConfig struct {
	Host string
	Port string
}

func InitConfig() {
	dc := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	rc := RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	hc := HTTPConfig{
		Port: os.Getenv("HTTP_PORT"),
	}

	config = Config{DB: dc, Redis: rc, Http: hc}
}
