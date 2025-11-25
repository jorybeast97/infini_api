package handler

import (
    "net/http"
    "infini_api/src/config"
    "infini_api/src/dao/memory"
    "infini_api/src/server"
)

// Handler is the Vercel Serverless Function entry point
func Handler(w http.ResponseWriter, r *http.Request) {
    cfg := config.LoadConfig()
    store := memory.NewMemoryStore()
    memory.Seed(store)
    repos := server.BuildRepos(store)
    services := server.BuildServices(cfg, repos)
    engine := server.BuildRouter(cfg, services)
    engine.ServeHTTP(w, r)
}

