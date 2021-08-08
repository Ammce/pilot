package entity

type Person struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"min=20" validate:"is-adult"`
	Email     string `json:"email"`
}
type Video struct {
	Id          string `json:"id"`
	Title       string `json:"title" binding:"max=20"`
	Description string `json:"description"`
	URL         string `json:"url" validate:"required"`
	Author      Person `json:"author" binding:"required"`
}
