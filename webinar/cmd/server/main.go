package main

import (
	"fmt"
	"net/http"
	"os"
	"yandex_practicum/webinar/internal/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	controller := controllers.NewBaseController(log)

	r := chi.NewRouter()
	r.Mount("/", controller.Route())

	log.Info("Server started")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Printf("Server:" + err.Error())
	}
}
