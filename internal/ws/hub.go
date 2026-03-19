package ws

import (
	"encoding/json"
	"log"
	"sync"
)

// Hub holds active WebSocket connections and broadcasts messages (e.g. streak updates, notifications).
type Hub struct {
	clients    map[string]map[*Client]bool // userID -> set of clients
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

// NewHub creates a new WebSocket hub.
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		broadcast:  make(chan []byte, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the hub loop (must be called in a goroutine).
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.mu.Lock()
			if h.clients[c.userID] == nil {
				h.clients[c.userID] = make(map[*Client]bool)
			}
			h.clients[c.userID][c] = true
			h.mu.Unlock()

		case c := <-h.unregister:
			h.mu.Lock()
			if m, ok := h.clients[c.userID]; ok {
				delete(m, c)
				if len(m) == 0 {
					delete(h.clients, c.userID)
				}
			}
			close(c.send)
			h.mu.Unlock()

		case msg := <-h.broadcast:
			h.mu.RLock()
			for _, m := range h.clients {
				for c := range m {
					select {
					case c.send <- msg:
					default:
						close(c.send)
						delete(m, c)
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

// SendToUser sends a message to all connections for the given user (e.g. for notifications).
func (h *Hub) SendToUser(userID string, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("ws SendToUser marshal: %v", err)
		return
	}
	h.mu.RLock()
	defer h.mu.RUnlock()
	if m, ok := h.clients[userID]; ok {
		for c := range m {
			select {
			case c.send <- data:
			default:
				// skip full buffer
			}
		}
	}
}
