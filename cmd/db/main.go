package main

import (
	"context"
	"fmt"

	"github.com/Wilsonator123/Learn/config"
)
func main() {
	conn, err := config.New()
	if err != nil {
		fmt.Println("Failed to create connection:", err)
		return
	}
	if conn == nil {
		fmt.Println("Connection is nil")
		return
	}
	
	err = conn.QueryRow(context.Background(), "CREATE DATABASE progress").Scan()

	if err != nil {
		fmt.Println("Failed to create Database")
	}
	fmt.Println("Created Database progress")
}