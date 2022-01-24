package main

import (
	"net/http"
	"time"

	"ticket-wallet/application/service"
	"ticket-wallet/presentation/controller"
	"ticket-wallet/presentation/router"

	"github.com/prometheus/common/log"
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
