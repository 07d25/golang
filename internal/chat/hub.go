// Package chat implementa el hub concurrente de las salas WebSocket.
package chat

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/go-chi/chi/v5"
	"noniserver/m/v2/internal/auth"
)

type Event struct {
	Type    string `json:"type"`
	RoomID  string `json:"room_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
type Hub struct {
	mu        sync.RWMutex
	rooms     map[string]map[*Client]struct{}
	save      func(context.Context, Event) (Event, error)
	authorize func(context.Context, string, string) error
}
type Client struct {
	conn   *websocket.Conn
	send   chan Event
	hub    *Hub
	roomID string
	userID string
}

func NewHub(save func(context.Context, Event) (Event, error), authorize func(context.Context, string, string) error) *Hub {
	return &Hub{rooms: make(map[string]map[*Client]struct{}), save: save, authorize: authorize}
}
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomID")
	userID := ""
	if id, ok := auth.UserID(r.Context()); ok {
		userID = strconv.FormatInt(id, 10)
	}
	if h.authorize != nil {
		if err := h.authorize(r.Context(), roomID, userID); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
	}
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return
	}
	conn.SetReadLimit(1024)
	client := &Client{conn: conn, send: make(chan Event, 16), hub: h, roomID: roomID, userID: userID}
	h.add(client)
	go client.writeLoop(r.Context())
	client.readLoop(r.Context())
}
func (h *Hub) add(c *Client) {
	h.mu.Lock()
	if h.rooms[c.roomID] == nil {
		h.rooms[c.roomID] = make(map[*Client]struct{})
	}
	h.rooms[c.roomID][c] = struct{}{}
	h.mu.Unlock()
	h.broadcast(c.roomID, Event{Type: "user_joined", RoomID: c.roomID, UserID: c.userID})
}
func (h *Hub) remove(c *Client) {
	h.mu.Lock()
	delete(h.rooms[c.roomID], c)
	close(c.send)
	if len(h.rooms[c.roomID]) == 0 {
		delete(h.rooms, c.roomID)
	}
	h.mu.Unlock()
	h.broadcast(c.roomID, Event{Type: "user_left", RoomID: c.roomID, UserID: c.userID})
}
func (h *Hub) broadcast(roomID string, event Event) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for client := range h.rooms[roomID] {
		select {
		case client.send <- event:
		default:
		}
	}
}
func (c *Client) readLoop(ctx context.Context) {
	defer c.hub.remove(c)
	defer c.conn.Close(websocket.StatusNormalClosure, "closed")
	for {
		var event Event
		if err := wsjson.Read(ctx, c.conn, &event); err != nil {
			return
		}
		event.RoomID = c.roomID
		event.UserID = c.userID
		if c.hub.save != nil {
			saved, err := c.hub.save(ctx, event)
			if err != nil {
				_ = wsjson.Write(ctx, c.conn, Event{Type: "error", RoomID: c.roomID, Content: err.Error()})
				continue
			}
			event = saved
		}
		c.hub.broadcast(c.roomID, event)
	}
}
func (c *Client) writeLoop(ctx context.Context) {
	defer c.conn.Close(websocket.StatusNormalClosure, "closed")
	keepAlive := time.NewTicker(30 * time.Second)
	defer keepAlive.Stop()
	for {
		select {
		case event, ok := <-c.send:
			if !ok {
				return
			}
			if err := wsjson.Write(ctx, c.conn, event); err != nil {
				return
			}
		case <-ctx.Done():
			return
		case <-keepAlive.C:
			pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			if err := c.conn.Ping(pingCtx); err != nil {
				cancel()
				return
			}
			cancel()
		}
	}
}
