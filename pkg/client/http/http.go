package http

import (
	"encoding/json"
	"errors"

	"github.com/Zettablock/zetta-sdk/pkg/config"
	"github.com/go-resty/resty/v2"
)

type (
	Http interface {
		Post(string, string, map[string]string, map[string]interface{}) (map[string]interface{}, error)
	}
	httpImpl struct {
		c *resty.Client
		f *config.HttpConfig
	}
)

func New(c *config.Config) Http {
	return &httpImpl{
		c: resty.New(),
		f: c.Client.Http,
	}
}

func (h *httpImpl) r(caller string, headers map[string]string) (*resty.Request, error) {
	req := h.c.R()
	headers["Content-Type"] = "application/json"
	for k, v := range headers {
		req = req.SetHeader(k, v)
	}

	authc, ok := h.f.Auth[caller]
	if !ok {
		return nil, errors.New("auth not configured for caller")
	}
	switch authc.Type {
	case "basic":
		req = req.SetBasicAuth(authc.User, authc.Password)
	case "token":
		req = req.SetAuthToken(authc.Token)
	}

	return req, nil
}

func (h *httpImpl) Post(
	caller string,
	url string,
	headers map[string]string,
	body map[string]interface{},
) (map[string]interface{}, error) {
	req, err := h.r(caller, headers)
	if err != nil {
		return nil, err
	}

	resp, err := req.SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New("non-200 status code")
	}

	r := map[string]interface{}{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, errors.New("unmarshal error")
	}
	return r, nil
}
