// Command seed carga tres usuarios de demostración para probar la aplicación.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"noniserver/m/v2/internal/database"
)

type demoUser struct {
	username    string
	displayName string
	bio         string
	password    string
}

var users = []demoUser{
	{username: "pixel_ana", displayName: "Ana Pixel", bio: "Exploradora de mundos abiertos.", password: "pixel-ana-123"},
	{username: "ninja_luis", displayName: "Luis Ninja", bio: "Ranked, café y estrategia.", password: "ninja-luis-123"},
	{username: "retro_mia", displayName: "Mía Retro", bio: "Coleccionista de controles.", password: "retro-mia-123"},
}

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Printf("cargar .env: %v", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL no está configurada; copia .env.example como .env")
	}
	db, err := database.Open(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := database.Migrate(ctx, db); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.ExecContext(ctx, `
			INSERT INTO users (username, display_name, bio, password_hash, deleted_at)
			VALUES ($1, $2, $3, $4, NULL)
			ON CONFLICT (username) DO UPDATE SET
				display_name = EXCLUDED.display_name,
				bio = EXCLUDED.bio,
				password_hash = EXCLUDED.password_hash,
				deleted_at = NULL`, user.username, user.displayName, user.bio, string(hash))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Seed listo: 3 usuarios de demostración disponibles.")
	for _, user := range users {
		fmt.Printf("- %-12s password: %s\n", user.username, user.password)
	}
}
