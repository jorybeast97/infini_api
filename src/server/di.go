package server

import (
    "log"
    "infini_api/src/config"
    "infini_api/src/dao"
    "infini_api/src/dao/memory"
    "infini_api/src/dao/mysql"
    "infini_api/src/dao/postgres"
    authservice "infini_api/src/service/auth"
    authorsservice "infini_api/src/service/authors"
    postsservice "infini_api/src/service/posts"
    usersservice "infini_api/src/service/users"
)

type Repos struct {
    Authors dao.AuthorsRepository
    Posts   dao.PostsRepository
    Photos  dao.PhotosRepository
    Apps    dao.AppsRepository
    Users   dao.UsersRepository
}

type Services struct {
    Auth    *authservice.AuthServiceImpl
    Authors *authorsservice.AuthorsServiceImpl
    Posts   *postsservice.PostsServiceImpl
    Users   *usersservice.UsersServiceImpl
    Tokens  authservice.TokenService
}

func BuildRepos(store *memory.MemoryStore) Repos {
    if pdb, err := postgres.Connect(); err == nil {
        return Repos{
            Authors: mysql.NewAuthorsRepo(pdb),
            Posts:   mysql.NewPostsRepo(pdb),
            Photos:  mysql.NewPhotosRepo(pdb),
            Apps:    mysql.NewAppsRepo(pdb),
            Users:   mysql.NewUsersRepo(pdb),
        }
    }
    db, err := mysql.Connect()
    if err == nil {
        if err := mysql.AutoMigrate(db); err != nil { log.Printf("migrate error: %v", err) }
        return Repos{
            Authors: mysql.NewAuthorsRepo(db),
            Posts:   mysql.NewPostsRepo(db),
            Photos:  mysql.NewPhotosRepo(db),
            Apps:    mysql.NewAppsRepo(db),
            Users:   mysql.NewUsersRepo(db),
        }
    }
    return Repos{
        Authors: memory.NewAuthorsRepo(store),
        Posts:   memory.NewPostsRepo(store),
        Photos:  memory.NewPhotosRepo(store),
        Apps:    memory.NewAppsRepo(store),
        Users:   memory.NewUsersRepo(store),
    }
}

func BuildServices(cfg config.Config, r Repos) Services {
    jwt := authservice.NewJWTService(cfg.Secret)
    return Services{
        Auth:    authservice.NewAuthService(jwt, r.Users),
        Authors: authorsservice.NewAuthorsService(r.Authors),
        Posts:   postsservice.NewPostsService(r.Posts),
        Users:   usersservice.NewUsersService(r.Users),
        Tokens:  jwt,
    }
}
