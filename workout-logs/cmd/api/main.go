package main

import (
	"cloud-gym/internal/adapter"
	"cloud-gym/internal/config"
	"cloud-gym/internal/mongo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	// swagger
	_ "cloud-gym/docs"
	"github.com/swaggo/http-swagger"
)

// @title Smart-Gym API
// @version 1.0
// @description This is the strongest server ever.

// @BasePath /api/v1
// @accept json
func main() {
	mongo.GetConnection()
	mqttClient := adapter.NewMosquittoClient(
		config.MosquittoDomain,
		config.MqttPort,
		config.MqttClientId,
		config.MqttUsername,
		config.MqttPassword,
		config.MqttCleanSession == "true",
	)

	defer (*mqttClient).Disconnect(1000)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Mount("/api/v1", adapter.MakeAppRouter())

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3030/swagger/doc.json"), //The url pointing to API definition
	))

	log.Println("Server is running on port 3030")
	http.ListenAndServe(":3030", r)
}
