package helper

import (
	"github.com/joho/godotenv"
	"os"
)

func ReadENV(envFile, envVar string) string {
	if envFile == "" {
		envFile = ".env"
	}
	var err error
	if err = godotenv.Load(envFile); err != nil {
		panic(err.Error())
	}
	return os.Getenv(envVar)
}
