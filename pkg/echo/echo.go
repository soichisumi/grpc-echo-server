package echo

import (
	"context"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

type EchoServer struct{}

func (EchoServer) Echo(ctx context.Context, req *grpctesting.EchoRequest) (*grpctesting.EchoResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	logger.Info("", zap.Any("req", req), zap.Any("md", md))
	return &grpctesting.EchoResponse{
		Message:              req.Message,
	}, nil
}

func (EchoServer) Empty(ctx context.Context, req *grpctesting.EmptyRequest) (*grpctesting.EmptyResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	logger.Info("", zap.Any("req", req), zap.Any("md", md))
	return &grpctesting.EmptyResponse{}, nil
}
