package health

import (
	"github.com/soichisumi/protodep/logger"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewHealthServer() *HealthServer {return &HealthServer{}}

type HealthServer struct {}

func (h HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	logger.Info("health check")
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h HealthServer) Watch(req *grpc_health_v1.HealthCheckRequest, s grpc_health_v1.Health_WatchServer) error {
	logger.Info("health check")
	if err := s.Send(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING }); err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return err
	}
	return nil
}
