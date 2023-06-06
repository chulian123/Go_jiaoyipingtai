package config

import (
	"jobcenter/internal/database"
	"jobcenter/internal/kline"
)

type Config struct {
	Okx   kline.OkxConfig
	Mongo database.MongoConfig
}
