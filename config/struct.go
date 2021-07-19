/**
 * @Author: Resynz
 * @Date: 2021/7/19 13:52
 */
package config

import (
	"golang.org/x/net/websocket"
	"sync"
)

type PlatformType uint8

const (
	PlatformTypeUnknown PlatformType = iota
	PlatformPC
	PlatformH5
	PlatformIOS
	PlatformAndroid
)

type logConf struct {
	Path         string `json:"path"`
	Name         string `json:"name"`
	Level        string `json:"level"`
	RequestIdKey string `json:"request_id_key"`
}

type Config struct {
	Mode      string  `json:"mode"`
	AppPort   int     `json:"app_port"`
	LogConfig logConf `json:"log_config"`
	AuthUrl   string  `json:"auth_url"`
}

type Client struct {
	Conn       *websocket.Conn
	UserId     string
	ClientId   string
	Platform   PlatformType
	CreateTime int64
}

type SafeClientMap struct {
	sync.RWMutex
	clientMap map[string][]*Client
}

func (s *SafeClientMap) Read(key string) []*Client {
	s.RLock()
	defer s.RUnlock()
	return s.clientMap[key]
}

func (s *SafeClientMap) Write(key string, clients []*Client) {
	s.Lock()
	defer s.Unlock()
	s.clientMap[key] = clients
}

func (s *SafeClientMap) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.clientMap, key)
}

func (s *SafeClientMap) Exists(key string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.clientMap[key]
	return ok
}

func (s *SafeClientMap) Size() int {
	return len(s.clientMap)
}

func (s *SafeClientMap) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range s.clientMap {
		keys = append(keys, k)
	}
	return keys
}
