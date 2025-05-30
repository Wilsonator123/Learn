package helper

import (
	"slices"

	"github.com/Wilsonator123/Learn/repository"
)

type GroupedColumns struct {
	Position int16 `json:"position"`
	Tasks []repository.List `json:"tasks"`
}
func GroupTasksByColumn(tasks []repository.List) []GroupedColumns {
	response := []GroupedColumns{}
	var columnSeen []int16;

	for _, task := range tasks {
		var column int16
		if !task.Position.Valid {
			column = 1
		} else {
			column = task.Position.Int16
		}
		
		idx := slices.IndexFunc(columnSeen, func(c int16) bool { return c == column })

		if idx == -1 {
			response = append(response, GroupedColumns{
				Position: column,
				Tasks: []repository.List{task},
			})
		} else {
			response[idx] = GroupedColumns{
				Position: response[idx].Position,
				Tasks: append(response[idx].Tasks, task),
			}
		}
	}

	return response
}