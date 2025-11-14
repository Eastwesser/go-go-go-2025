package config

import (
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	MaxConcurrentRequests int           `env:"MAX_CONCURRENT_REQUESTS" envDefault:"10"`
	RateLimitPerSecond    int           `env:"RATE_LIMIT_PER_SECOND" envDefault:"100"`
	CacheTTL              time.Duration `env:"CACHE_TTL" envDefault:"5m"`
	NumWorkers            int           `env:"NUM_WORKERS" envDefault:"5"`
	MaxRetries            int           `env:"MAX_RETRIES" envDefault:"3"`
	JSONDataPath          string        `env:"JSON_DATA_PATH" envDefault:"reyna_route.json"`
	DebugMode             bool          `env:"DEBUG_MODE" envDefault:"false"`
	ServerPort            string        `env:"SERVER_PORT" envDefault:"8080"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	
	return &cfg, nil
}