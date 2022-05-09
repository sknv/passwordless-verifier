package main

import (
	"flag"
	stdlog "log"
)

func main() {
	configPath := ConfigFilePathFlag()
	flag.Parse()

	cfg, err := ParseConfig(*configPath)
	if err != nil {
		stdlog.Fatalf("parse config: %s", err)
	}

	stdlog.Print(cfg)
}
