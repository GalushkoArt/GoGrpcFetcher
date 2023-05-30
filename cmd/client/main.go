package main

import (
	"context"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/config"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/logs"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/fetcher"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/service"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.Init()
	logs.Init(config.Conf.Logs.Level, "client.txt")

	client, err := service.NewFetcherClient("localhost:50051")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create client!")
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	runClient(ctx, client)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)
	<-exit

	done := make(chan bool)
	go func() {
		cancelFunc()
		err := client.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close client!")
		}
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

func runClient(ctx context.Context, client *service.FetcherClient) {
	incrementCounter := int64(0)
	for {
		select {
		case <-time.After(1 * time.Minute):
			response, err := client.SendFetchRequest(ctx, &fetcher.FetchRequest{Url: "http://164.92.251.245:8080/api/v1/products/"})
			if err != nil {
				log.Err(err).Msg("Failed to send fetch request!")
			} else {
				log.Info().Msgf("Response: %v", response)
			}
		case <-time.After(10 * time.Second):
			response, err := client.SendListRequest(ctx, &fetcher.GetRequest{Sorting: []*fetcher.SortParams{{Field: fetcher.SortField_PRICE, Asc: false}}, Paging: &fetcher.PagingParams{Page: incrementCounter, PageSize: 100}})
			incrementCounter++
			if err != nil {
				log.Err(err).Msg("Failed to send list request!")
			} else {
				log.Info().Msgf("Response: %v", response)
			}
		case <-ctx.Done():
			return
		}
	}
}
