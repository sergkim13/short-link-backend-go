package model

type Link struct {
	Id 			int 	`json:"id"`
	Original 	string 	`json:"original"`
	Short 		string 	`json:"short"`
	Clicks 		int 	`json:"clicks"`
	CreatedAt 	string 	`json:"created_at"`
	UpdatedAt 	string 	`json:"updated_at"`
}
