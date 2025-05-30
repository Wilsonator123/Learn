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
	Tasks []repository.Task `json:"tasks"`
	Position int16 `json:"position"`
}

func ListAll() ([]helper.GroupedColumns, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return []helper.GroupedColumns{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	tasks, err := queries.GetAllTasks(ctx)

	response := helper.GroupTasksByColumn(tasks)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return []helper.GroupedColumns{}, err
	}

	conn.Close(ctx)
	
	return response, nil
}

func GetTask(id string) (repository.Task, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return repository.Task{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	parsedUUID, err := uuid.Parse(id)
	
	if err != nil {
		fmt.Printf("Failed to parse UUID: %v\n", err)
		return repository.Task{}, err
	}

	response, err := queries.GetTask(ctx, parsedUUID)
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
		return repository.Task{}, err
	}

	conn.Close(ctx)

	return response, nil
}

func CreateTask(input model.NewTask) (repository.Task, error) {
	ctx := context.Background()
	conn, err := config.New()

	if err != nil {
		return repository.Task{}, errors.New("database connection failed")
	}

	queries := repository.New(conn)

	id := uuid.New()

	now := time.Now()

	newTask := repository.CreateNewTaskParams{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if input.Position != nil {
		newTask.Position = pgtype.Int2{
			Int16: *input.Position,
			Valid:  true,
		}
	}
	

	task, err := queries.CreateNewTask(ctx, newTask)

	if err != nil {
		fmt.Printf("Failed to create user %v\n", err)
		return repository.Task{}, err
	}

	conn.Close(ctx)

	return task, nil
}

func DeleteTask(id string) bool {
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

	err = queries.DeleteTask(ctx, parsedUUID)
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
