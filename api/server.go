package handler

import (
	"infini_api/src/config"
	"infini_api/src/dao/memory"
	"infini_api/src/dao/postgres"
	"infini_api/src/server"
	"net/http"
	"os"
	"sync"
)

var (
	initOnce sync.Once
	engine   http.Handler
)

// Handler is the Vercel Serverless Function entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	initOnce.Do(func() {
		cfg := config.LoadConfig()
		store := memory.NewMemoryStore()
		repos := server.BuildRepos(store)
		services := server.BuildServices(cfg, repos)
		ginEngine := server.BuildRouter(cfg, services)
		// Optional one-time migration controlled by env
		if os.Getenv("MIGRATE") == "1" {
			if db, err := postgres.Connect(); err == nil {
				_ = postgres.AutoMigrate(db)
			}
		}
		engine = ginEngine
	})
	engine.ServeHTTP(w, r)
}
