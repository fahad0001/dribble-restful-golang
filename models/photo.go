package models

type Photo struct {
	Id int	`json:"id" db:"id"`
	Title string	`json:"title" db:"title"`
	Description string	`json:"description" db:"description"`
	FileName string	`json:"file-name" db:"file_name"`
	DribbleLink imageType	`json:"images"`
}

type imageType struct {
	Normal string `json:"normal" db:"dribble_link"`
}