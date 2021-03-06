package server

import (
	"time"
)

type Config struct {
	MaxUnzippedSize int64         `config:"max_unzipped_size"`
	MaxHeaderBytes  int           `config:"max_header_bytes"`
	ReadTimeout     time.Duration `config:"read_timeout"`
	WriteTimeout    time.Duration `config:"write_timeout"`
	SecretToken     string        `config:"secret_token"`
	SSL             *SSLConfig    `config:"ssl"`
}

type SSLConfig struct {
	Enabled    *bool  `config:"enabled"`
	PrivateKey string `config:"key"`
	Cert       string `config:"certificate"`
}

func (c *SSLConfig) IsEnabled() bool {
	return c != nil && (c.Enabled == nil || *c.Enabled)
}

var defaultConfig = Config{
	MaxUnzippedSize: 10 * 1024 * 1024, // 10mb
	MaxHeaderBytes:  1048576,          // 1mb
	ReadTimeout:     2 * time.Second,
	WriteTimeout:    2 * time.Second,
	SecretToken:     "",
}
