package model


type LinkCreate struct {
	OriginalURL string 	`binding:"required" json:"original_url"`  //nolint:tagliatelle  //linter is wrong here
}
type Link struct {
	ID 			int 	`json:"id"`
	OriginalURL string 	`db:"original"   json:"original"`
	ShortURL 	string 	`db:"short"      json:"short"`
	CreatedAt 	string 	`db:"created_at" json:"created_at"` 	 //nolint:tagliatelle  //linter is wrong here
}
