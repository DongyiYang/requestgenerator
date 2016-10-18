package generator

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
)

type RequestGenerator interface {
	GenerateRequest(address string) error
}

type RealRequestGenerator struct {
	duration int64
	qps      float64

	counter int64
}

func NewRealRequestGenerator(duration int64, qps float64) *RealRequestGenerator {
	return &RealRequestGenerator{
		duration,
		qps,
		0,
	}
}

func (this *RealRequestGenerator) GenerateRequest(address string) error {
	startTimestamp := int64(time.Now().Unix())
	for this.duration == -1 || int64(time.Now().Unix())-startTimestamp <= this.duration {
		err := this.sendRequest(address)
		if err != nil {
			return fmt.Errorf("Error sending request to %s: %v", address, err)
		}

		this.counter = this.counter + 1

		err = this.wait()
		if err != nil {
			return fmt.Errorf("Error sending request periodically: %v", err)
		}
	}
	glog.V(3).Infof("Finished!")
	return nil
}

func (this *RealRequestGenerator) sendRequest(address string) error {
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	glog.V(3).Infof("Request #%d has been sent to %s", this.counter, address)
	return nil
}

// Time off between two request.
func (this *RealRequestGenerator) wait() error {
	var timeBase time.Duration
	var base float64
	switch {
	case this.qps < 1:
		timeBase = time.Second
		base = 1
	case this.qps < 1000:
		timeBase = time.Millisecond
		base = 1000
	case this.qps < 1E6:
		timeBase = time.Microsecond
		base = 1E6
	default:
		return fmt.Errorf("Invalid QPS value: %d.", this.qps)
	}

	time.Sleep(time.Duration(int32(base/this.qps)) * timeBase)

	return nil
}
