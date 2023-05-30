package service

import (
	"context"
	"encoding/csv"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

type FetcherService interface {
	Fetch(ctx context.Context, url string) (model.FetchStatus, error)
	ListItems(ctx context.Context, options model.GetOptions) ([]model.Item, error)
}

type FetcherRepository interface {
	StartStatus(ctx context.Context, url string) (*primitive.ObjectID, error)
	FinishStatus(ctx context.Context, id *primitive.ObjectID) error
	GetStatus(ctx context.Context, url string) (*model.FetchStatus, error)
}

type ItemUpdateRepository interface {
	UpsertItem(ctx context.Context, items []model.Item) error
	GetItems(ctx context.Context, options model.GetOptions) ([]model.Item, error)
}

type fetcherService struct {
	fetcherRepository FetcherRepository
	itemRepository    ItemUpdateRepository
	client            http.Client
}

func NewFetcherService(clientTimeout time.Duration, fetcherRepository FetcherRepository, itemRepository ItemUpdateRepository) FetcherService {
	return &fetcherService{
		fetcherRepository: fetcherRepository,
		itemRepository:    itemRepository,
		client: http.Client{
			Timeout: clientTimeout,
		},
	}
}

func (s *fetcherService) Fetch(ctx context.Context, url string) (model.FetchStatus, error) {
	fetchStatus, err := s.fetcherRepository.GetStatus(ctx, url)
	if fetchStatus != nil && fetchStatus.FetchStatus == model.IN_PROGRESS {
		return model.FetchStatus{FetchStatus: model.IN_PROGRESS, Started: fetchStatus.Started}, nil
	}
	started := time.Now()
	status, err := s.fetcherRepository.StartStatus(ctx, url)
	if err != nil {
		return model.FetchStatus{FetchStatus: model.FAILED, Started: started}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return model.FetchStatus{FetchStatus: model.FAILED, Started: started}, err
	}
	response, err := s.client.Do(req)
	if err != nil || response.StatusCode != 200 {
		return model.FetchStatus{FetchStatus: model.FAILED, Started: started}, err
	}
	reader := csv.NewReader(response.Body)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return model.FetchStatus{FetchStatus: model.FAILED, Started: started}, err
	}
	items, err := recordsToItems(records)
	if err != nil {
		return model.FetchStatus{FetchStatus: model.FAILED, Started: started}, err
	}
	err = s.itemRepository.UpsertItem(ctx, items)
	if err != nil {
		return model.FetchStatus{FetchStatus: model.FAILED, Started: started}, err
	}
	err = s.fetcherRepository.FinishStatus(ctx, status)
	return model.FetchStatus{FetchStatus: model.DONE, Started: started, Finished: time.Now()}, nil
}

func recordsToItems(records [][]string) ([]model.Item, error) {
	items := make([]model.Item, 0, len(records))
	for _, record := range records {
		price, err := strconv.Atoi(record[1])
		if err != nil {
			log.Error().Err(err).Msg("Failed to convert price to int! Value: " + record[1])
			return nil, err
		}
		items = append(items, model.Item{
			Name:  record[0],
			Price: price,
		})
	}
	return items, nil
}

func (s *fetcherService) ListItems(ctx context.Context, options model.GetOptions) ([]model.Item, error) {
	return s.itemRepository.GetItems(ctx, options)
}
