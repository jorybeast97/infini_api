package mysql

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "infini_api/src/domain"
    "fmt"
    "os"
)

func Connect() (*gorm.DB, error) {
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    name := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    if user == "" { user = "root" }
    if pass == "" { pass = "19970911Fh" }
    if name == "" { name = "infini" }
    if host == "" { host = "127.0.0.1" }
    if port == "" { port = "3306" }
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &domain.Author{},
        &domain.BlogPost{},
        &domain.Photo{},
        &domain.AppProject{},
        &domain.User{},
    )
}
