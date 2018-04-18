package storage

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/pkg/errors"

	"github.com/bobrofon/back-pocket/internal/app/back-pocket/config"
)

type ProxyStorage struct {
	client  *http.Client
	cache   map[string]string
	rwMutex *sync.RWMutex
	selfURL *url.URL
}

func NewProxyStorage(conf *config.BackPocketConf) *ProxyStorage {
	transport := &http.Transport{Proxy: http.ProxyURL(conf.HTTPProxy)}
	client := &http.Client{Transport: transport}
	selfURL := &url.URL{Scheme: "http", Host: conf.ListenAddr}
	rwMutex := &sync.RWMutex{}

	return &ProxyStorage{client: client, selfURL: selfURL, rwMutex: rwMutex}
}

func (s *ProxyStorage) Put(key string, value string) error {
	s.setCache(key, value)
	_, err := s.Get(key)
	if err != nil {
		return errors.Wrap(err, "cannot get key via proxy")
	}
	return nil
}

func (s *ProxyStorage) Get(key string) (string, error) {
	getURL := *s.selfURL
	getURL.Path = key
	req, err := http.NewRequest(http.MethodGet, getURL.RequestURI(), nil)
	if err != nil {
		return "", errors.Wrap(err, "cannot construct get request")
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "cannot execute get request")
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("unexpected proxy response status: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "cannot read response body")
	}
	value := string(body)
	s.setCache(key, value)
	return value, nil
}

func (s *ProxyStorage) getCache(key string) string {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.cache[key]
}

func (s *ProxyStorage) setCache(key string, value string) {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.cache[key] = value
}
