package bnrenv

import (
	"log"
	"os"
	"path/filepath"
	"time"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func Load(cfg interface{}, filename ...string) {

	if err := godotenv.Load(filename...); err != nil {
		log.Fatal(".env read failed ", err)
	}

	if err := env.Parse(cfg); err != nil {
		log.Fatal("parse .env failed ", err)
	}
}

func GetLocalPath(path string) string {
	localPath := filepath.Join(path, time.Now().Format("2006-01-02"))
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		_ = os.MkdirAll(localPath, os.ModePerm)
	}
	return localPath
}
