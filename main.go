package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/echo"
	"github.com/soichisumi/grpc-echo-server/pkg/health"
	grpctesting "github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	defaultPort = 8080
)

func main(){
	var port = flag.Int("p", defaultPort, "port number for listening")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	//creds, err := credentials.NewServerTLSFromFile("./certs/cert.pem", "./certs/privkey.pem")
	//if err != nil {
	//	logger.Fatal(err.Error(), zap.Error(err))
	//}
	//server := grpc.NewServer(grpc.Creds(creds))

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
				md, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					logger.Debug("request metadata is empty")
				} else {
					logger.Info("request metadata", zap.Any("metadata", md))
				}
				return handler(ctx, req)
			},
		),
	)
	grpctesting.RegisterEchoServiceServer(server, echo.NewEchoServer())
	grpc_health_v1.RegisterHealthServer(server, health.NewHealthServer())
	reflection.Register(server)

	logger.Info("", zap.Int("port", *port))
	if err := server.Serve(lis); err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
}
