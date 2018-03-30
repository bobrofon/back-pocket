package config

import (
	"net/url"

	"github.com/pkg/errors"
)

type BackPocketConf struct {
	HttpProxy *url.URL
}

func NewBackPocketConf(rawUrl string) (*BackPocketConf, error) {
	if httpProxy, err := url.Parse(rawUrl); err != nil {
		return nil, errors.Wrap(err, "invalid proxy url")
	} else {
		return &BackPocketConf{HttpProxy:httpProxy}, nil
	}
}
