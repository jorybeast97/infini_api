package main

import (
	"infini_api/src/config"
	"infini_api/src/dao/memory"
	"infini_api/src/server"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	store := memory.NewMemoryStore()
	memory.Seed(store)
	repos := server.BuildRepos(store)
	services := server.BuildServices(cfg, repos)
	engine := server.BuildRouter(cfg, services)

	addr := ":" + cfg.Port
	log.Printf("Infini API listening on %s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatal(err)
	}
}
