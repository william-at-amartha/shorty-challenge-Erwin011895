package main

import (
	"github.com/Erwin011895/shorty-challenge/cmd/webservice"
	"github.com/Erwin011895/shorty-challenge/internal/config"
)

func main() {
	config.Init()
	webservice.Start()
}

