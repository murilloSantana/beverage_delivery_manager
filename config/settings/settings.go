package settings

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"time"
)

type OptSettings func(*Settings)

type Settings struct {
	Port          string
	MongoSettings MongoSettings
	RedisSettings RedisSettings
}

type MongoSettings struct {
	DatabaseName    string
	CollectionName  string
	URL             string
	MinPoolSize     uint64
	MaxPoolSize     uint64
	MaxConnIdleTime time.Duration
}

type RedisSettings struct {
	URL          string
	Database     int
	Password     string
	MinIdleConns int
	IdleTimeout  time.Duration
	PoolSize     int
}

func New(opts ...OptSettings) Settings {
	s := &Settings{
		Port:          os.Getenv("PORT"),
		MongoSettings: newMongoSettings(),
		RedisSettings: newRedisSettings(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return *s
}

func newMongoSettings() MongoSettings {
	maxIdle, _ := strconv.Atoi(os.Getenv("MONGO_MAX_CONN_IDLE_TIME"))

	maxConnIdleTime := time.Minute * time.Duration(maxIdle)
	minPool, _ := strconv.Atoi(os.Getenv("MONGO_MIN_POOL_SIZE"))
	maxPool, _ := strconv.Atoi(os.Getenv("MONGO_MAX_POOL_SIZE"))

	return MongoSettings{
		DatabaseName:    os.Getenv("MONGO_DB_NAME"),
		CollectionName:  os.Getenv("MONGO_COLLECTION_NAME"),
		URL:             os.Getenv("MONGO_URL"),
		MinPoolSize:     uint64(minPool),
		MaxPoolSize:     uint64(maxPool),
		MaxConnIdleTime: maxConnIdleTime,
	}
}

func newRedisSettings() RedisSettings {
	idleTimeout := 5 * time.Minute
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	poolSize, _ := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	minIdleConns, _ := strconv.Atoi(os.Getenv("REDIS_MIN_IDLE_CONN"))

	return RedisSettings{
		URL:          os.Getenv("REDIS_URL"),
		Database:     db,
		Password:     os.Getenv("REDIS_PASSWORD"),
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		IdleTimeout:  idleTimeout,
	}
}
