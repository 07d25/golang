package chat

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/go-chi/chi/v5"
)

func TestHubBroadcastsEventToRoom(t *testing.T) {
	hub := NewHub(nil, nil)
	first := &Client{send: make(chan Event, 1), hub: hub, roomID: "1"}
	second := &Client{send: make(chan Event, 1), hub: hub, roomID: "1"}
	hub.add(first)
	hub.add(second)
	<-first.send
	<-second.send
	hub.broadcast("1", Event{Type: "message", RoomID: "1", Content: "hola"})
	for name, client := range map[string]*Client{"first": first, "second": second} {
		select {
		case event := <-client.send:
			if event.Content != "hola" {
				t.Fatalf("%s recibió un evento inesperado: %+v", name, event)
			}
		default:
			t.Fatalf("%s no recibió el evento", name)
		}
	}
}

func TestHubBroadcastsAcrossWebSocketConnections(t *testing.T) {
	hub := NewHub(nil, nil)
	router := chi.NewRouter()
	router.Handle("/ws/rooms/{roomID}", hub)
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Skipf("entorno sin sockets locales: %v", err)
	}
	server := &http.Server{Handler: router}
	go server.Serve(listener)
	defer server.Shutdown(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := "ws://" + listener.Addr().String() + "/ws/rooms/1"
	first, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer first.Close(websocket.StatusNormalClosure, "done")
	second, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer second.Close(websocket.StatusNormalClosure, "done")

	// Consumimos los eventos de presencia generados durante las conexiones.
	readEvent(t, ctx, first)
	readEvent(t, ctx, first)
	readEvent(t, ctx, second)
	if err := wsjson.Write(ctx, first, Event{Type: "message", Content: "hola"}); err != nil {
		t.Fatal(err)
	}
	event := readEvent(t, ctx, second)
	if event.Type != "message" || event.Content != "hola" || event.RoomID != "1" {
		t.Fatalf("evento inesperado: %+v", event)
	}
}

func readEvent(t *testing.T, ctx context.Context, conn *websocket.Conn) Event {
	t.Helper()
	var event Event
	if err := wsjson.Read(ctx, conn, &event); err != nil {
		t.Fatal(err)
	}
	return event
}
