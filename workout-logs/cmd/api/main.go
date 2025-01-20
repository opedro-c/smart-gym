package main

import (
	"cloud-gym/internal/adapter"
	"cloud-gym/internal/config"
	"cloud-gym/internal/mongo"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	// swagger
	_ "cloud-gym/docs"

	"github.com/swaggo/http-swagger"
)

//	@title			Smart-Gym API
//	@version		1.0
//	@description	This is the strongest server ever.

// @BasePath	/api/v1
// @accept		json
func main() {
	setUpLogger()
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
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Mount("/api/v1", adapter.MakeAppRouter())

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3030/swagger/doc.json"), //The url pointing to API definition
	))

	slog.Info("Server started at :3030")
	http.ListenAndServe(":3030", r)
}

func setUpLogger() {
	logLevelEnv := config.LogLevel
	var logLevel slog.Level
	switch logLevelEnv {
	case "DEBUG":
		slog.Info("Setting log level to DEBUG")
		logLevel = slog.LevelDebug
	case "INFO":
		slog.Info("Setting log level to INFO")
		logLevel = slog.LevelInfo
	case "ERROR":
		slog.Info("Setting log level to ERROR")
		logLevel = slog.LevelError
	}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
