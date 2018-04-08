package config

import (
	"github.com/bobrofon/back-pocket/internal/app/back-pocket/constant"
	"testing"
)

func TestNewBackPocketConf(t *testing.T) {
	validUrls := []string{
		"http://127.0.0.1",
	}
	invalidUrls := []string{
		":8888",
	}
	validAddr := ":5555"

	for _, u := range validUrls {
		if _, err := NewBackPocketConf(u, validAddr); err != nil {
			t.Error(
				"For:", u,
				"expected: valid",
				"got:", err,
			)
		}
	}

	for _, u := range invalidUrls {
		if conf, err := NewBackPocketConf(u, validAddr); err == nil {
			t.Error(
				"For:", u,
				"expected: invalid",
				"got:", conf,
			)
		}
	}

	if conf, err := NewBackPocketConf(validUrls[0], ""); err == nil {
		t.Error("For empty listen address expected: invalid got:", conf)
	}
}

func TestCurrentBackPocketConf(t *testing.T) {
	proxy := constant.DefaultHTTPProxy
	bind := constant.DefaultBindAddress
	conf, err := CurrentBackPocketConf()
	if err != nil {
		t.Error(
			"For proxy:", proxy, "bind:", bind,
			"expected: valid",
			"got:", err,
		)
	}
	if conf.ListenAddr != bind {
		t.Error(
			"For proxy:", proxy, "bind:", bind,
			"result bind:", conf.ListenAddr,
		)
	}
	if proxy != conf.HTTPProxy.String() {
		t.Error(
			"For proxy:", proxy, "bind:", bind,
			"result proxy:", conf.HTTPProxy,
		)
	}
}
