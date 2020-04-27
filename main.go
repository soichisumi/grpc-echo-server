package main

import (
	"flag"
	"fmt"
	"github.com/soichisumi/go-util/logger"
	"github.com/soichisumi/grpc-echo-server/pkg/echo"
	"github.com/soichisumi/grpc-echo-server/pkg/health"
	grpctesting "github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

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

	server := grpc.NewServer()
	grpctesting.RegisterEchoServiceServer(server, echo.NewEchoServer())
	grpc_health_v1.RegisterHealthServer(server, health.NewHealthServer())

	logger.Info("", zap.Int("port", *port))
	if err := server.Serve(lis); err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
}
