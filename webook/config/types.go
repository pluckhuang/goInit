package config

import "time"

type config struct {
	DB         DBConfig
	Redis      RedisConfig
	LocalCache LocalCacheConfig
}

type DBConfig struct {
	DSN string
}
type RedisConfig struct {
	Addr string
}

type LocalCacheConfig struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}
