package settings

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type OptSettings func(*Settings)

type Settings struct {
	Port string
}

func New(opts ...OptSettings) Settings {
	s := &Settings{
		Port: os.Getenv("PORT"),
	}

	for _, opt := range opts {
		opt(s)
	}

	return *s
}
