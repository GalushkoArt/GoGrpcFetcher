package gRPC

import (
	"fmt"
	"github.com/GalushkoArt/GoGrpcFetcher/internal/service"
	"github.com/GalushkoArt/GoGrpcFetcher/pkg/fetcher"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type Server struct {
	grpcSrv        *grpc.Server
	fetcherHandler *fetcherHandler
}

var grpcLog zerolog.Logger

func StartGRPC(service service.FetcherService, port int) func() {
	grpcLog = log.With().Str("from", "grpcServer").Logger()
	return New(service).ListenAndServe(port)
}

func New(service service.FetcherService) *Server {
	kaep := keepalive.EnforcementPolicy{
		MinTime:             10 * time.Second,
		PermitWithoutStream: true,
	}
	kasp := keepalive.ServerParameters{
		Time:    5 * time.Minute,
		Timeout: 5 * time.Minute,
	}
	return &Server{
		grpcSrv: grpc.NewServer(
			grpc.KeepaliveEnforcementPolicy(kaep),
			grpc.KeepaliveParams(kasp),
		),
		fetcherHandler: newFetcherHandler(service),
	}
}

func (s *Server) ListenAndServe(port int) func() {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		grpcLog.Panic().Err(err).Msgf("Failed to listen port :%d", port)
	}
	fetcher.RegisterFetcherServiceServer(s.grpcSrv, s.fetcherHandler)

	go func() {
		if err := s.grpcSrv.Serve(lis); err != nil {
			grpcLog.Info().Err(err).Msg("gRPC server has stopped")
		}
	}()

	grpcLog.Info().Msgf("gRPC server started successfully on port %s", addr)
	return s.grpcSrv.GracefulStop
}
