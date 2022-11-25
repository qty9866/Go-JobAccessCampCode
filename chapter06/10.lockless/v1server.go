package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Content string
}

type WebServerV1 struct {
	config     Config
	configLock sync.RWMutex
}

func (ws *WebServerV1) Reload() {
	ws.configLock.Lock()
	defer ws.configLock.Unlock()
	ws.config.Content = fmt.Sprintf("%v", time.Now().UnixNano())
}

func (ws *WebServerV1) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		ws.Reload()
	}
}

func (ws *WebServerV1) visit() string {
	ws.configLock.RLock()
	defer ws.configLock.RUnlock()
	return ws.config.Content
}
