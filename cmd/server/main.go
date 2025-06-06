package main

import (
	"gofootball/internal/config"
	"gofootball/internal/server"
)

func main() {
	cfg := config.GetConfig()

	srv := &server.Server{}
	srv.Initialize(cfg)
	srv.Run(":3000")
}
