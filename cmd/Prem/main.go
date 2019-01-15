package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"

	"Prem/pkg/api"
)

const httpPort int = 8080

func main() {
	initLogrus()
	router := api.Routes()

	corsHandler := handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))
	logrus.WithFields(logrus.Fields{
		"http-port": httpPort,
	}).Info("Starting server")
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", httpPort), corsHandler(&router)))

}

type formatter struct {
	lf logrus.Formatter
}

func (f *formatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	_, e.Caller.File = filepath.Split(e.Caller.File)
	return f.lf.Format(e)
}

func initLogrus() {
	logrus.SetFormatter(&formatter{&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000000+00:00",
	}})
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
}
