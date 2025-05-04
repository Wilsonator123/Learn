package model

type NewItem struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Priority    *int16 `json:"priority,omitempty"`
}