package models

type Item struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
