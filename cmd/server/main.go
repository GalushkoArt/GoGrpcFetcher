package main

import (
	"context"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/config"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/logs"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/repository"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/service"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/transport/gRPC"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.Init()
	logs.Init(config.Conf.Logs.Level, config.Conf.Logs.Path)
	opts := options.Client()
	mongoConf := config.Conf.Mongo
	opts.SetAuth(options.Credential{
		Username: mongoConf.Username,
		Password: mongoConf.Password,
	})
	opts.ApplyURI(mongoConf.URI)
	dbClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}
	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping MongoDB")
	}

	db := dbClient.Database(config.Conf.Mongo.Database)
	log.Info().Msgf("Connected to %s mongo database successfully!", mongoConf.Database)

	mongoRepository := repository.NewMongoRepository(db)
	fetcherService := service.NewFetcherService(15*time.Second, mongoRepository, mongoRepository)
	grpcShutdown := gRPC.StartGRPC(fetcherService, config.Conf.GRPC.Port)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)
	<-exit

	done := make(chan bool)
	go func() {
		grpcShutdown()
		done <- true
	}()
	select {
	case <-time.After(2 * time.Minute):
		log.Error().Msg("Failed to shutdown in 2 Minutes")
		os.Exit(1)
	case <-done:
		log.Info().Msg("Shutdown successfully")
	}
}
