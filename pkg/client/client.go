package client

import (
	"github.com/Zettablock/zetta-sdk/pkg/client/http"
	"github.com/Zettablock/zetta-sdk/pkg/client/logger"
	"github.com/Zettablock/zetta-sdk/pkg/config"
	"go.uber.org/zap"
)

type Client struct {
	Config *config.Config
	Logger *zap.SugaredLogger
	Http   http.Http
}

func New(c *config.Config) (*Client, error) {
	l, err := logger.New(c)
	if err != nil {
		return nil, err
	}

	h := http.New(c)
	return &Client{
		Config: c,
		Logger: l,
		Http:   h,
	}, nil
}
