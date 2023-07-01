package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRCPClientConn(cfg *TransportConfig) (conn *grpc.ClientConn, err error) {
	return grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			// propagate header, retry and sending external metrics count
			propagateAndObservationUnaryClientInterceptor(cfg),
		),
	)
}
