package postgres

import (
	"infini_api/src/domain"
	"io"
	stdlog "log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	url := os.Getenv("DB_URL")
	if url == "" {
		return nil, os.ErrNotExist
	}
	cfg := &gorm.Config{
		Logger: logger.New(
			stdlog.New(io.Discard, "", stdlog.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent,
				Colorful:      false,
			},
		),
	}
	return gorm.Open(postgres.Open(url), cfg)
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
