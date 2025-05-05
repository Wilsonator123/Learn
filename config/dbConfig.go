package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)
func New() (*pgx.Conn, error) {
	godotenv.Load()
	value := os.Getenv("DATABASE_URL")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, value)
	if err != nil {
		fmt.Printf("Unable to connect to database %v\n", err)
	}

	return conn, err
}