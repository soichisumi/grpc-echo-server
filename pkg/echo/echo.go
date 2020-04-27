package echo

import (
	"context"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.uber.org/zap"
)

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

type EchoServer struct{}

func (EchoServer) Echo(ctx context.Context, req *grpctesting.EchoRequest) (*grpctesting.EchoResponse, error) {
	logger.Info("", zap.Any("req", req))
	return &grpctesting.EchoResponse{
		Message:              req.Message,
	}, nil
}

func (EchoServer) Empty(ctx context.Context, req *grpctesting.EmptyRequest) (*grpctesting.EmptyResponse, error) {
	logger.Info("", zap.Any("req", req))
	return &grpctesting.EmptyResponse{}, nil
}
