package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

type (
	Config struct {
		Client *ClientConfig `json:"client"`
	}

	ClientConfig struct {
		Logger *LoggerConfig `json:"logger"`
		Http   *HttpConfig   `json:"http"`
	}

	LoggerConfig struct {
		Level    string `json:"level"`
		Encoding string `json:"encoding"`
	}

	HttpConfig struct {
		Auth map[string]*HttpAuthConfig `json:"auth"`
	}

	HttpAuthConfig struct {
		Type     string `json:"type"`
		Token    string `json:"token"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)

func New() func() (*Config, error) {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l.Sugar()

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return func() (*Config, error) {
		viper.SetConfigName("base")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(path.Join(filepath.Dir(exe), "../pkg/config/yaml"))
		logger.Infow("creating config")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Errorw("failed to read in config", "err", err)
			return nil, err
		}

		config := &Config{}
		err = viper.Unmarshal(config)
		if err != nil {
			logger.Errorw("failed to unmarshal config", "err", err)
			return nil, err
		}

		logger.Infow("service config is built", "config", config)
		return config, nil
	}
}
