package main

import (
	"net/http"
	"time"

	"github.com/vsyakunin/ticket-wallet/application/service"
	"github.com/vsyakunin/ticket-wallet/presentation/controller"
	"github.com/vsyakunin/ticket-wallet/presentation/router"

	log "github.com/sirupsen/logrus"
)

const httpTimeout = 30 * time.Second

func main() {
	svc := service.NewService()
	cont := controller.NewController(svc)

	router := router.NewRouter(cont)
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  httpTimeout,
		WriteTimeout: httpTimeout,
	}

	log.Info("Server is up and running")
	log.Fatal(server.ListenAndServe().Error())
}
