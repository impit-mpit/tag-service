package main

import (
	"neuro-most/tags-service/config"
	"neuro-most/tags-service/internal/infra"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	infra.Config(cfg).Database().Serve().Start()
}
