package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"url-shortener/pkg/api"
	"url-shortener/pkg/config"
	"url-shortener/pkg/dal"

	"github.com/sirupsen/logrus"
)

const (
	allowConnectionsAfterShutdown = 5 * time.Second
)

func main() {
	// application config
	cfg := config.New()
	logrus.Info("ENV: ok")

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.InfoLevel)

	// dal
	dal, err := dal.New(cfg)
	if err != nil {
		logrus.Fatalf("Unable to setup dal, Error: %v ", err.Error())
	}
	logrus.Info("DAL: ok")

	// construct the api and run it
	a := &api.API{
		Config: cfg,
		DAL:    dal,
	}

	go func() {
		log.Fatal(a.Serve())
	}()

	// graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan

	logrus.Infof("Request to shutdown server. Doing nothing for %v", allowConnectionsAfterShutdown)
	waitTimer := time.NewTimer(allowConnectionsAfterShutdown)
	<-waitTimer.C

	logrus.Infof("Shutting down server...")
	logrus.Fatal(a.Shutdown())
}
