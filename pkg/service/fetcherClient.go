package service

import (
	"context"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/fetcher"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FetcherClient struct {
	client fetcher.FetcherServiceClient
	conn   *grpc.ClientConn
}

func NewFetcherClient(address string) (*FetcherClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := fetcher.NewFetcherServiceClient(conn)
	log.Info().Msg("Fetcher client connected!")
	return &FetcherClient{client: client, conn: conn}, nil
}

func (c *FetcherClient) Close() error {
	return c.conn.Close()
}

func (c *FetcherClient) SendFetchRequest(ctx context.Context, request *fetcher.FetchRequest) (*fetcher.FetchResponse, error) {
	return c.client.Fetch(ctx, request)
}

func (c *FetcherClient) SendListRequest(ctx context.Context, request *fetcher.GetRequest) (*fetcher.GetResponse, error) {
	return c.client.List(ctx, request)
}
