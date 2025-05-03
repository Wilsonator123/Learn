package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Wilsonator123/Learn/config"
	"github.com/Wilsonator123/Learn/model"
	"github.com/Wilsonator123/Learn/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PriorityList struct{
	NoPriority []repository.List
	Other []repository.List
}

func ListAll() (PriorityList, error) {
	ctx := context.Background()
	conn := config.New()
	queries := repository.New(conn)
	var response PriorityList;

	rows, err := queries.GetAllItems(ctx)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return PriorityList{}, err
	}
	
	for i := range rows {
		priority, _ := pgtype.Text.MarshalJSON(rows[i].Priority)
		if(string(priority) != "null" ) {
			response.Other = append(response.Other, rows[i])
		} else {
			response.NoPriority = append(response.NoPriority, rows[i])
		}
	}

	conn.Close(ctx)
	
	return response, nil
}

func GetItem(id string) (repository.List, error) {
	ctx := context.Background()
	conn := config.New()
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
	conn := config.New()
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

	if input.Priority != nil {
		newItem.Priority = pgtype.Text{
			String: *input.Priority,
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
	conn := config.New()
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
