package config

import "testing"

func TestNewBackPocketConf(t *testing.T) {
	validUrls := []string{
		"http://127.0.0.1",
	}
	invalidUrls := []string{
		":8888",
	}

	for _, u := range validUrls {
		if _, err := NewBackPocketConf(u); err != nil {
			t.Error(
				"For:", u,
				"expected: valid",
				"got:", err,
			)
		}
	}

	for _, u := range invalidUrls {
		if conf, err := NewBackPocketConf(u); err == nil {
			t.Error(
				"For:", u,
				"expected: invalid",
				"got:", conf,
			)
		}
	}
}
