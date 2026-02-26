package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ArdhiCode/go-auth/cmd"
	"github.com/joho/godotenv"
)

func initEnv() {
	if _, err := os.Stat(".env"); err == nil {
		path, _ := filepath.Abs(".env")
		if err := godotenv.Load(path); err != nil {
			log.Printf("Warning : failed to load env at %s : %v", path, err)
		} else {
			log.Printf("Success to load env at %s", path)
			purl := os.Getenv("POSTGRES_URL")
			if purl != "" {
				log.Printf("POSTGRES_URL length: %d", len(purl))
			} else {
				log.Printf("POSTGRES_URL is EMPTY!")
			}
		}
	}
}

func main() {
	initEnv()

	if err := cmd.Commands(); err != nil {
		panic("Failed get commands: " + err.Error())
	}
}
