package community

import (
	"context"
	"strings"
	"testing"
)

func TestCreatePostRejectsInvalidContentBeforeDatabase(t *testing.T) {
	store := &Store{}
	if _, err := store.CreatePost(context.Background(), 1, postRequest{}); err == nil {
		t.Fatal("se esperaba error para contenido vacío")
	}
	if _, err := store.CreatePost(context.Background(), 1, postRequest{Content: strings.Repeat("x", 501)}); err == nil {
		t.Fatal("se esperaba error para contenido demasiado largo")
	}
}

func TestCreateMessageRejectsInvalidInputBeforeDatabase(t *testing.T) {
	store := &Store{}
	if _, err := store.CreateMessage(context.Background(), 1, messageRequest{}); err == nil {
		t.Fatal("se esperaba error para mensaje inválido")
	}
	if _, err := store.CreateMessage(context.Background(), 1, messageRequest{UserID: 1, Content: strings.Repeat("x", 501)}); err == nil {
		t.Fatal("se esperaba error para mensaje demasiado largo")
	}
}
