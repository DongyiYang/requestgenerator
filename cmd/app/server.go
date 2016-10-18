package app

import (
	"fmt"
	"net"

	"github.com/dongyiyang/requestgenerator/cmd/app/options"
	"github.com/dongyiyang/requestgenerator/pkg/generator"
)

type RequestGeneratorServer struct {
	config *options.RequestGeneratorConfig
}

func NewRequestGeneratorServer(config *options.RequestGeneratorConfig) (*RequestGeneratorServer, error) {
	if config.Host == "" {
		return nil, fmt.Errorf("Host of remote is not set.")
	}

	if config.QPS < 1 {
		return nil, fmt.Errorf("Invalid QPS: %d", config.QPS)
	}

	return &RequestGeneratorServer{config}, nil
}

func (s RequestGeneratorServer) Run() error {
	address := s.config.Host
	if s.config.Port != "" {
		address = net.JoinHostPort(address, s.config.Port)
	}
	fmt.Printf("Duration is %d\n", s.config.Duration)
	g := generator.NewRealRequestGenerator(s.config.Duration, s.config.QPS)
	return g.GenerateRequest(address)
}
