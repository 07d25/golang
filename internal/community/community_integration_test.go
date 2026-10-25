package community

import (
	"context"
	"os"
	"testing"

	"noniserver/m/v2/internal/database"
)

func TestCommunityFlowWithPostgres(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL no configurada")
	}
	ctx := context.Background()
	db, err := database.Open(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := database.Migrate(ctx, db); err != nil {
		t.Fatal(err)
	}
	_, _ = db.ExecContext(ctx, "TRUNCATE messages, room_members, rooms, comments, likes, posts, follows, users RESTART IDENTITY CASCADE")

	store := NewStore(db)
	ana, err := store.CreateUser(ctx, userRequest{Username: "integration_ana", DisplayName: "Ana", Password: "integration-password"})
	if err != nil {
		t.Fatal(err)
	}
	room, err := store.CreateRoom(ctx, roomRequest{Name: "integration-room"})
	if err != nil {
		t.Fatal(err)
	}
	if err := store.JoinRoom(ctx, room.ID, ana.ID); err != nil {
		t.Fatal(err)
	}
	message, err := store.CreateMessage(ctx, room.ID, messageRequest{UserID: ana.ID, Content: "mensaje persistido"})
	if err != nil {
		t.Fatal(err)
	}
	messages, err := store.ListMessages(ctx, room.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(messages) != 1 || messages[0].ID != message.ID || messages[0].UserID != ana.ID {
		t.Fatalf("historial inesperado: %+v", messages)
	}
}
