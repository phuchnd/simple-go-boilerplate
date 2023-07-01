package grpc

type TransportConfig struct {
	ServiceName         string
	ExternalServiceName string

	Host            string
	Port            int
	MaxRetries      int
	BackoffDelaysMs int
}
