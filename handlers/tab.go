package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Wilsonator123/Learn/config"
	"github.com/Wilsonator123/Learn/repository"
	"github.com/google/uuid"
)

type Tabs struct {
	Tabs []repository.Tab `json:"tabs"`
}

func ListTabs() ([]repository.Tab, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return []repository.Tab{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	tabs, err := queries.GetAllTabs(ctx)

	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return []repository.Tab{}, err
	}

	conn.Close(ctx)

	return tabs, nil
}

func CreateTab() error {
	ctx := context.Background()
	conn, err := config.New()

	

	id := uuid.New()
	now := time.Now()

	newTab := repository.CreateNewTabParams{
		ID: id,
		Title: "Tab",
		Layout: json.RawMessage("{}"),
		CreatedAt: now,
		UpdatedAt: now,
	};

	if err != nil {
		return errors.New("database connection failed")
	}

	queries := repository.New(conn)

	_, err = queries.CreateNewTab(ctx, newTab)

	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return err
	}

	conn.Close(ctx)

	return nil
}
