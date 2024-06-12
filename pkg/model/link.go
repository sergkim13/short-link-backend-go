package model


type LinkCreate struct {
	OriginalURL string 	`json:"original_url" binding:"required"`
}
type Link struct {
	Id 			int 	`json:"-"`
	OriginalURL string 	`json:"original"`
	ShortURL 	string 	`json:"short"`
	CreatedAt 	string 	`json:"created_at"`
}
