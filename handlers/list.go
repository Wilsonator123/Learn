package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Wilsonator123/Learn/config"
	"github.com/Wilsonator123/Learn/helper"
	"github.com/Wilsonator123/Learn/model"
	"github.com/Wilsonator123/Learn/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Column struct {
	Tasks []repository.List `json:"tasks"`
	Position int16 `json:"position"`
}

func ListAll() ([]helper.GroupedColumns, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return []helper.GroupedColumns{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	tasks, err := queries.GetAllItems(ctx)

	response := helper.GroupTasksByColumn(tasks)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return []helper.GroupedColumns{}, err
	}

	conn.Close(ctx)
	
	return response, nil
}

func GetItem(id string) (repository.List, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return repository.List{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	parsedUUID, err := uuid.Parse(id)
	
	if err != nil {
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return repository.List{}, err
	}

	response, err := queries.GetItem(ctx, parsedUUID)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return repository.List{}, err
	}

	conn.Close(ctx)

	return response, nil
}

func CreateItem(input model.NewItem) (repository.List, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return repository.List{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	id := uuid.New()

	now := time.Now()

	newItem := repository.CreateNewItemParams{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if input.Position != nil {
		newItem.Position = pgtype.Int2{
			Int16: *input.Position,
			Valid:  true,
		}
	}
	

	item, err := queries.CreateNewItem(ctx, newItem)

	if err != nil {
		fmt.Printf("Failed to create user %v\n", err)
		return repository.List{}, err
	}

	conn.Close(ctx)

	return item, nil
}

func DeleteItem(id string) bool {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return false
	}

	queries := repository.New(conn)

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return false
	}

	err = queries.DeleteItem(ctx, parsedUUID)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return false
	}

	return true
	
}

// func UpdateItem(id string, new_item repository.List) bool {
// 	ctx := context.Background()
// 	conn := config.New()
// 	queries := repository.New(conn)

// 	parsedUUID, err := uuid.Parse(id)
// 	if err != nil {
// 		fmt.Printf("Failed to parse UUID: %v\n", err)
// 		return false
// 	}

// 	old_item, err := queries.GetItem(ctx, parsedUUID)
// 	if err != nil {
// 		fmt.Printf("Failed with error: %v\n", err)
// 		return false
// 	}

// 	err = queries.UpdateItem(ctx, repository.UpdateItemParams{
// 		...old_item
// 	})
	

// 	return true
// }
