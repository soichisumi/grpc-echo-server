package main

import (
	"context"
	"crypto/tls"
	"github.com/soichisumi/go-util/logger"
	grpctesting "github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	endpoint = "xxx.xxx.xxx.xxx"
	port = "443"
)

func main(){
	//conn, err := grpc.Dial(endpoint+":"+port, grpc.WithInsecure())
	conn, err := grpc.Dial(endpoint+":"+port, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	defer conn.Close()

	c := grpctesting.NewEchoServiceClient(conn)
	res, err := c.Echo(context.Background(), &grpctesting.EchoRequest{
		Message: "yo, test request",
	})
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}
	logger.Info("", zap.Any("res", res))
}