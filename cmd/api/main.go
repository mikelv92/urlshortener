package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/mikelv92/urlshortener/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	Logger *slog.Logger
	URLMappingModel models.URLMappingModelInterface
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	mongoClient, err := openMongoCLient()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	app := Application{
		Logger: logger,
		URLMappingModel: &models.URLMappingModel{MongoClient: mongoClient},
	}

	server := http.Server{
		Addr: ":4000",
		Handler: app.routes(),
		ErrorLog: slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting server")
	err = server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openMongoCLient() (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/urls"))
	if err != nil {
		return nil, err
	}

	return mongoClient, nil
}