package dto

type CreateItemInput struct {
	Title string `json:"title" binding:"required,min=2"`
}

type UpdateItemInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}
