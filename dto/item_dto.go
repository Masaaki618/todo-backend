package dto

type CreateItemInput struct {
	Title string `json:"title" binding:"required,min=2"`
}
