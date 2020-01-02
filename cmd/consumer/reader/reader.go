package reader

import (
	"github.com/consumer-firebase-token/internal/cache"
	"github.com/consumer-firebase-token/internal/config"
	"github.com/consumer-firebase-token/internal/consumer"
	"github.com/consumer-firebase-token/internal/db"
)

// Reader holds all the data relevant.
type Reader struct {
	DB       *db.DB
	Consumer *consumer.Consumer
	Cache    *cache.Cache
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	c := consumer.NewConsumer(cfg)

	ch, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &Reader{
		DB:       dbs,
		Consumer: c,
		Cache:    ch,
	}, nil
}
