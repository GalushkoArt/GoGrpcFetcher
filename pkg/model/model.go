package model

import (
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/fetcher"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FetchStatus struct {
	ID          primitive.ObjectID `json:"_" bson:"_id,omitempty"`
	URL         string             `json:"URL" bson:"url,omitempty"`
	FetchStatus Status             `json:"FetchStatus" bson:"fetch_status,omitempty"`
	Started     time.Time          `json:"Started,omitempty" bson:"started,omitempty"`
	Finished    time.Time          `json:"Finished,omitempty" bson:"finished,omitempty"`
}

type Status string

const (
	DONE        Status = "DONE"
	IN_PROGRESS Status = "IN_PROGRESS"
	FAILED      Status = "FAILED"
)

func (s Status) ToGrpcStatus() fetcher.FetchStatus {
	switch s {
	case DONE:
		return fetcher.FetchStatus_DONE
	case IN_PROGRESS:
		return fetcher.FetchStatus_IN_PROGRESS
	case FAILED:
		return fetcher.FetchStatus_FAILED
	default:
		log.Fatal().Str("from", "model").Msgf("Unknown status %v", s)
		return fetcher.FetchStatus_FAILED
	}
}

type Item struct {
	Name  string             `json:"name" bson:"name,omitempty"`
	Price int                `json:"price" bson:"price,omitempty"`
	From  primitive.ObjectID `json:"_" bson:"from,omitempty"`
}

type GetOptions struct {
	NameSort  *int
	PriceSort *int
	Paging    *Paging
}

type Paging struct {
	Page     int64
	PageSize int64
}
