package config

import (
	"net/url"

	"github.com/pkg/errors"
	backend "github.com/spf13/viper"

	"github.com/bobrofon/back-pocket/internal/app/back-pocket/constant"
)

type BackPocketConf struct {
	HTTPProxy  *url.URL
	ListenAddr string
}

func NewBackPocketConf(proxyURL string, listenAddr string) (*BackPocketConf, error) {
	if len(listenAddr) == 0 {
		return nil, errors.New("empty listed address is not allowed")
	}
	httpProxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, errors.Wrap(err, "invalid proxy url")
	}
	return &BackPocketConf{
		HTTPProxy:  httpProxy,
		ListenAddr: listenAddr,
	}, nil
}

func CurrentBackPocketConf() (*BackPocketConf, error) {
	return NewBackPocketConf(backend.GetString(constant.HTTPProxy),
		backend.GetString(constant.BindAddress))
}
