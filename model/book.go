package model

type Book struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
