package handlers

import (
	"context"
	"fmt"

	"github.com/Wilsonator123/Learn/config"
	"github.com/Wilsonator123/Learn/repository"
	"github.com/google/uuid"
)


func ListAll() []repository.List {
	ctx := context.Background()
	conn := config.New()

	queries := repository.New(conn)

	response, err := queries.GetAllItems(ctx)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
	}
	
	return response
}

func GetItem(id string) repository.List {
	ctx := context.Background()
	conn := config.New()
	queries := repository.New(conn)

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return repository.List{}
	}

	response, err := queries.GetItem(ctx, parsedUUID)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
	}

	return response
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
