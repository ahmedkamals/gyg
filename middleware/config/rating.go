package config

type (
	Item struct {
		Id int `json:"tour_id"`
		Rating	string	`json:"rating"`
	}

	Ratings []Item
)
