package main

import (
	"github.com/sirupsen/logrus"

	"Prem/pkg/api"
	"sync"
)

const httpPort int = 8080

func main() {
	apiServer := api.NewServer(httpPort)
	logrus.WithFields(logrus.Fields{
		"port": httpPort,
	}).Info("Starting HTTP server")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err := apiServer.Run()
		if err != nil {
			logrus.Fatal(err)
		}
		wg.Done()
	}()
	defer apiServer.Close()

	wg.Wait()
}
