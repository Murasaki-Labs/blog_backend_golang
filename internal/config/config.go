package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	BindPort       int           `env:"SRV_HTTP_PORT,default=8080"`
	BindHost       string        `env:"SRV_BIND_ADDRESS,default=0.0.0.0"`
	MetricsPort    int           `env:"SRV_METRICS_PORT,default=9000"`
	DebugMod       bool          `env:"DEBUG_MOD"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT,default=20s"`
	LogFormat      string        `env:"LOG_FORMAT,default=json"`
	LogLevel       string        `env:"LOG_LEVEL,default=info"`
	LogName        string        `env:"LOG_NAME,default=blog-backend"`
	LogType        string        `env:"LOG_TYPE,default=local"`
}

func FromEnv(ctx context.Context) (*Config, error) {
	var c Config

	if err := envconfig.Process(ctx, &c); err != nil {
		return nil, fmt.Errorf("fatal error occurred while loading config: %w", err)
	}

	return &c, nil
}
