package config

import (
	"fmt"
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

// TLS/SSL Configs
type TLSConfig struct {
	UseTLS   bool   `yaml:"use_tls"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

// HTTPServer Configs
type HTTPServer struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
	TLS          TLSConfig     `yaml:"tls"`
}

// Configurations
type Config struct {
	Env     string     `yaml:"env"`
	Server  HTTPServer `yaml:"server"`
	Storage string     `yaml:"storage"`
}

func LoadConfig(configFile string) (Config, error) {
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("[LoadConfig] (%v)", err)
	}
	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("[LoadConfig] (%v)", err)
	}
	return config, nil
}
