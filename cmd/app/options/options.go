package options

import (
	"github.com/spf13/pflag"
)

type RequestGeneratorConfig struct {
	Host     string
	Port     string
	Duration int64
	QPS      float64
}

func NewRequestGeneratorConfig() *RequestGeneratorConfig {
	return &RequestGeneratorConfig{
		Duration: -1,
	}
}

func (s *RequestGeneratorConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Host, "host", s.Host, "The address of the remote server.")
	fs.StringVar(&s.Port, "port", s.Port, "The port number of the remote server. Optional.")
	fs.Int64Var(&s.Duration, "duration", s.Duration, "The time this generator should run. If not set, it will never exit.")
	fs.Float64Var(&s.QPS, "qps", s.QPS, "The frequency of the request. Request per second.")
}
