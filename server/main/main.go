package main

import (
	"flag"
	"github.com/kallydev/privacy/service"
	"log"
)

var configPath string

func init() {
	flag.Parse()
	flag.StringVar(&configPath, "config", "../config.yaml", "config file path")
}

func main() {
	svc := service.NewService(configPath)
	if err := svc.Start(); err != nil {
		log.Fatalln(err)
	}
}
