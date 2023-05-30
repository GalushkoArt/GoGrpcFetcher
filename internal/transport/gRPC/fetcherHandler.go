package gRPC

import (
	"context"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/service"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/fetcher"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fetcherHandler struct {
	service service.FetcherService
	fetcher.UnimplementedFetcherServiceServer
}

var fhLog zerolog.Logger

func newFetcherHandler(service service.FetcherService) *fetcherHandler {
	fhLog = log.With().Str("from", "fetcherHandler").Logger()
	return &fetcherHandler{service: service}
}

func (h *fetcherHandler) Fetch(ctx context.Context, request *fetcher.FetchRequest) (*fetcher.FetchResponse, error) {
	fetch, err := h.service.Fetch(ctx, request.Url)
	if err != nil {
		fhLog.Error().Err(err).Msg("Failed to fetch!")
		return &fetcher.FetchResponse{Status: fetch.FetchStatus.ToGrpcStatus(), Content: err.Error()}, status.Error(codes.Internal, err.Error())
	}
	return &fetcher.FetchResponse{Status: fetch.FetchStatus.ToGrpcStatus()}, nil
}

func (h *fetcherHandler) List(ctx context.Context, request *fetcher.GetRequest) (*fetcher.GetResponse, error) {
	options := model.GetOptions{}
	if request.Paging != nil {
		options.Paging = &model.Paging{
			Page:     request.Paging.Page,
			PageSize: request.Paging.PageSize,
		}
	}
	for _, sortParams := range request.Sorting {
		sortValue := 1
		if sortParams.Field == fetcher.SortField_NAME {
			if sortParams.Asc {
				sortValue = 1
			} else {
				sortValue = -1
			}
			options.NameSort = &sortValue
		} else if sortParams.Field == fetcher.SortField_PRICE {
			if sortParams.Asc {
				sortValue = 1
			} else {
				sortValue = -1
			}
			options.PriceSort = &sortValue
		}
	}
	list, err := h.service.ListItems(ctx, options)
	if err != nil {
		fhLog.Error().Err(err).Msg("Failed to get list of items!")
		return &fetcher.GetResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &fetcher.GetResponse{Items: mapItemsToGrpc(list), Sorting: request.Sorting, Paging: request.Paging}, nil
}

func mapItemsToGrpc(items []model.Item) []*fetcher.Item {
	result := make([]*fetcher.Item, 0, len(items))
	for _, item := range items {
		result = append(result, &fetcher.Item{Name: item.Name, Price: int32(item.Price)})
	}
	return result
}
