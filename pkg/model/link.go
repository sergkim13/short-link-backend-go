package model


type LinkCreate struct {
	OriginalURL string 	`json:"original_url" binding:"required"`
}
type Link struct {
	ID 			int 	`json:"id"`
	OriginalURL string 	`json:"original" db:"original"`
	ShortURL 	string 	`json:"short" db:"short"`
	CreatedAt 	string 	`json:"created_at" db:"created_at"`
}
