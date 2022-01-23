package main

import (
	"log"
	"net/http"
	"time"

	"ticket-wallet/application/service"
	"ticket-wallet/presentation/controller"
	"ticket-wallet/presentation/router"
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

	log.Println("Server is up and running")
	log.Fatal(server.ListenAndServe().Error())
}
