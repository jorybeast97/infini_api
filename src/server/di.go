package server

import (
    "infini_api/src/config"
    "infini_api/src/dao"
    "infini_api/src/dao/memory"
    authservice "infini_api/src/service/auth"
    authorsservice "infini_api/src/service/authors"
    postsservice "infini_api/src/service/posts"
)

type Repos struct {
    Authors dao.AuthorsRepository
    Posts   dao.PostsRepository
    Photos  dao.PhotosRepository
    Apps    dao.AppsRepository
}

type Services struct {
    Auth    *authservice.AuthServiceImpl
    Authors *authorsservice.AuthorsServiceImpl
    Posts   *postsservice.PostsServiceImpl
}

func BuildRepos(store *memory.MemoryStore) Repos {
    return Repos{
        Authors: memory.NewAuthorsRepo(store),
        Posts:   memory.NewPostsRepo(store),
        Photos:  memory.NewPhotosRepo(store),
        Apps:    memory.NewAppsRepo(store),
    }
}

func BuildServices(cfg config.Config, r Repos) Services {
    jwt := authservice.NewJWTService(cfg.Secret)
    return Services{
        Auth:    authservice.NewAuthService(jwt),
        Authors: authorsservice.NewAuthorsService(r.Authors),
        Posts:   postsservice.NewPostsService(r.Posts),
    }
}
