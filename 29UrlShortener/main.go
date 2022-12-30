package main

import (
	"os"

	"github.com/decipher07/Go-Getting/29UrlShortener/model"
	log "github.com/sirupsen/logrus"
)

/* Set up logs configs */
func setLogsConfig() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	setLogsConfig()

	/* Database Setup */
	model.Setup()
}
