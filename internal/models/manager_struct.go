package models

import (
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{}
)

type Manager struct {
	client ClientList
	sync.RWMutex
	otps RetentionMap
	handlers map[string]EventHandler
}

func NewManager(ctx context.Context) *Manager {
	return &Manager{
		client: make(ClientList),
		handlers: make(map[string]EventHandler),
		otps: NewRetentionMap(ctx, 5*time.Second),
	}
}

