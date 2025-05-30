package model

type NewTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Position    *int16 `json:"position,omitempty"`
}