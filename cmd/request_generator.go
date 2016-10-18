package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dongyiyang/requestgenerator/cmd/app"
	"github.com/dongyiyang/requestgenerator/cmd/app/options"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

func init() {
	flag.Set("logtostderr", "true")

}

func main() {
	config := options.NewRequestGeneratorConfig()
	config.AddFlags(pflag.CommandLine)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	s, err := app.NewRequestGeneratorServer(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	err = s.Run()
	if err != nil {
		glog.Fatalf("Error generating request: %v", err)
	}
}
